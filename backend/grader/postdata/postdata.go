package postdata

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"grader/auth"
	"grader/constant"
	"grader/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CreateRequestData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Isadmin   bool   `json:"isadmin"`
}

func CreateUser(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var createData CreateRequestData

	err = ctx.BindJSON(&createData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err := database.Connection.Exec("INSERT INTO account (Username, Password, Firstname, Lastname, Isadmin) VALUES (?, ?, ?, ?, ?)",
			createData.Username,
			createData.Password,
			createData.Firstname,
			createData.Lastname,
			createData.Isadmin,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while creating user",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "created user " + createData.Username,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type RequestUpdateGroupData struct {
	Groupname  string `json:"groupname"`
	Assignment gin.H  `json:"assignment"`
}

func UpdateGroupAssignment(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestUpdateGroupData

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		jsonbyte, err := json.Marshal(&updateData.Assignment)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while stringify data",
			})
			fmt.Println(err.Error())
			return
		}

		_, err = database.Connection.Exec("UPDATE grouplist SET assignment = ? WHERE groupname = ?",
			string(jsonbyte),
			updateData.Groupname,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating group data",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "updated data for group " + updateData.Groupname,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type TestcaseData struct {
	Name    string `json:"name"`
	Stdin   bool   `json:"stdin"`
	Action  string `json:"action"`
	Outfile string `json:"outfile"`
	Infile  string `json:"infile"`
}

type RequestTestcaseData struct {
	Id       string         `json:"id"`
	Testcase []TestcaseData `json:"testcase"`
}

type ResponseTestcaseData struct {
	Name  string `json:"name"`
	Stdin bool   `json:"stdin"`
}

func UpdateTestcase(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestTestcaseData

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		var directory string = constant.Directory + "/storage/sol/"
		folderPath := directory + updateData.Id

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			err := CreateFolder(folderPath)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error creating folder:" + err.Error(),
				})
				return
			}
		}

		if _, err := os.Stat(folderPath + "/in"); os.IsNotExist(err) {
			err := CreateFolder(folderPath + "/in")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error creating infolder:" + err.Error(),
				})
				return
			}
		}

		if _, err := os.Stat(folderPath + "/out"); os.IsNotExist(err) {
			err := CreateFolder(folderPath + "/out")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error creating outfolder:" + err.Error(),
				})
				return
			}
		}

		TestcaseDatas := []ResponseTestcaseData{}

		for _, v := range updateData.Testcase {

			var TestcaseData ResponseTestcaseData

			switch action := v.Action; action {

			case "add":

				TestcaseData.Name = v.Name
				TestcaseData.Stdin = v.Stdin
				if len(v.Outfile) > 0 {
					err := CreateUpdateFile(v.Outfile, folderPath+"/out/"+v.Name+".sol")
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error creating outfile for testcase:" + err.Error(),
						})
						return
					}
				}

				if v.Stdin && len(v.Infile) > 0 {
					err := CreateUpdateFile(v.Infile, folderPath+"/in/"+v.Name+".in")
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error creating infile for testcase:" + err.Error(),
						})
						return
					}
				}
				TestcaseDatas = append(TestcaseDatas, TestcaseData)
			case "remove":
				inPath := folderPath + "/in/" + v.Name + ".in"
				outPath := folderPath + "/out/" + v.Name + ".sol"

				if _, err := os.Stat(inPath); os.IsExist(err) {
					err := RemoveFile(inPath)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error deleting infile for testcase:" + err.Error(),
						})
						return
					}
				}
				if _, err := os.Stat(outPath); os.IsExist(err) {
					err := RemoveFile(outPath)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error deleting outfile for testcase:" + err.Error(),
						})
						return
					}
				}
			case "update":
				TestcaseData.Name = v.Name
				TestcaseData.Stdin = v.Stdin
				if len(v.Outfile) > 0 {
					err := CreateUpdateFile(v.Outfile, folderPath+"/out/"+v.Name+".sol")
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error creating outfile for testcase:" + err.Error(),
						})
						return
					}
				}

				if v.Stdin && len(v.Infile) > 0 {
					err := CreateUpdateFile(v.Infile, folderPath+"/in/"+v.Name+".in")
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"status":   "InternalError",
							"response": "Error creating infile for testcase:" + err.Error(),
						})
						return
					}
				}
				TestcaseDatas = append(TestcaseDatas, TestcaseData)
			default:
				TestcaseData.Name = v.Name
				TestcaseData.Stdin = v.Stdin
				TestcaseDatas = append(TestcaseDatas, TestcaseData)
			}

		}

		jsonbyte, err := json.Marshal(&TestcaseDatas)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while stringify data",
			})
			fmt.Println(err.Error())
			return
		}

		_, err = database.Connection.Exec("UPDATE assignment SET Testcase = ? WHERE Id = ?",
			string(jsonbyte),
			updateData.Id,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating testcase",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": TestcaseDatas,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type RequestTemplateData struct {
	Id     string `json:"id"`
	Action string `json:"action"`
	File   string `json:"file"`
}

func UpdateTemplate(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestTemplateData

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		var directory string = constant.Directory + "/storage/sol/"
		folderPath := directory + updateData.Id

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			err := CreateFolder(folderPath)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error creating folder:" + err.Error(),
				})
				return
			}
		}

		if _, err := os.Stat(folderPath + "/template"); os.IsNotExist(err) {
			err := CreateFolder(folderPath + "/template")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error creating infolder:" + err.Error(),
				})
				return
			}
		}
		switch action := updateData.Action; action {

		case "add":

			if len(updateData.File) > 0 {
				err := CreateUpdateFile(updateData.File, folderPath+"/template/"+updateData.Id+".tp")
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"status":   "InternalError",
						"response": "Error creating template file: " + err.Error(),
					})
					return
				}
			}
			_, err = database.Connection.Exec("UPDATE assignment SET Template = 1 WHERE Id = ?",
				updateData.Id,
			)

		case "remove":
			templatePath := folderPath + "/template/" + updateData.File + ".tp"
			if _, err := os.Stat(templatePath); os.IsExist(err) {
				err := RemoveFile(templatePath)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"status":   "InternalError",
						"response": "Error deleting template file : " + err.Error(),
					})
					return
				}
			}
			_, err = database.Connection.Exec("UPDATE assignment SET Template = 0 WHERE Id = ?",
				updateData.Id,
			)

		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating template data",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "Updated template",
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

func CreateFolder(path string) error {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return err
	}
	return nil
}

func CreateUpdateFile(filebase64 string, path string) error {

	fileBytes, err := base64.StdEncoding.DecodeString(filebase64)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, fileBytes, 0777)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return err
	}
	return nil
}

type RequestGroupData struct {
	Groupname string `json:"groupname"`
}

func CreateGroup(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestGroupData

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err = database.Connection.Exec("INSERT INTO grouplist (groupname) VALUES (?)",
			updateData.Groupname,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while creating user",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "created group " + updateData.Groupname,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

func RemoveGroup(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestGroupData

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err = database.Connection.Exec("DELETE FROM grouplist WHERE groupname=?",
			updateData.Groupname,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while creating user",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "removed group " + updateData.Groupname,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type RequestUpdateUserGroup struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}

func UpdateUserGroup(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestUpdateUserGroup

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		jsonbyte, err := json.Marshal(&updateData.Groups)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while stringify data",
			})
			fmt.Println(err.Error())
			return
		}

		_, err = database.Connection.Exec("UPDATE account SET groups = ? WHERE Username=?",
			string(jsonbyte),
			updateData.Username,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating user group",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "updated group " + updateData.Username,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type RequestCreateAssignment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
	Link string `json:"link"`
}

func CreateAssignment(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestCreateAssignment

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err := database.Connection.Exec("INSERT INTO assignment (Id, Name, Info, Link) VALUES (?, ?, ?, ?)",
			updateData.Id,
			updateData.Name,
			updateData.Info,
			updateData.Link,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while creating assignment",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "created assignement ",
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

func UpdateAssignment(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestCreateAssignment

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err := database.Connection.Exec("UPDATE assignment SET Name = ? , Info = ? ,Link = ? WHERE Id = ?",
			updateData.Name,
			updateData.Info,
			updateData.Link,
			updateData.Id,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating assignment",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "updated assignement ",
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type RequestUpdateGroupState struct {
	Groupname string `json:"groupname"`
	Active    bool   `json:"active"`
}

func UpdateGroupState(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := auth.ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := auth.IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var updateData RequestUpdateGroupState

	err = ctx.BindJSON(&updateData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "Data is not valid for server",
		})
		return
	}

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		_, err = database.Connection.Exec("UPDATE grouplist SET active = ? WHERE groupname = ?",
			updateData.Active,
			updateData.Groupname,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while updating group state",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": "updated group " + updateData.Groupname + " state",
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}
