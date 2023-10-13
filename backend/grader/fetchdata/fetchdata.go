package fetchdata

import (
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

type AdminTableStructure struct {
	Id       string
	Name     string
	Maxscore int
}

type Work struct {
	Id        string
	Pass      bool
	Timestamp int64
	Attempt   int
	Score     int
}

type AdmintableData struct {
	Username  string
	Firstname string
	Lastname  string
	Score     int
	Data      map[string]Work
}

type AdminTable struct {
	Structure map[string]AdminTableStructure
	Assigned  []AdmintableData
}

func GetUserTable(ctx *gin.Context) {
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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}
		result, err := database.Connection.Query(
			"SELECT Username,Firstname,Lastname,work FROM account",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		table := AdminTable{}

		for result.Next() {
			var unparseddata string
			var data AdmintableData
			err := result.Scan(&data.Username, &data.Firstname, &data.Lastname, &unparseddata)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			err = json.Unmarshal([]byte(unparseddata), &data.Data)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}
			data.Score = 0
			for _, v := range data.Data {
				if v.Pass {
					data.Score++
				}
			}

			table.Assigned = append(table.Assigned, data)
		}

		result, err = database.Connection.Query(
			"SELECT Id,Name,Testcase FROM assignment",
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching table struct data",
			})
			return
		}
		table.Structure = map[string]AdminTableStructure{}
		for result.Next() {
			var data AdminTableStructure
			var unparseddata string
			err := result.Scan(&data.Id, &data.Name, &unparseddata)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing structure data",
				})
				fmt.Println(err.Error())
				return
			}

			var Testcase []gin.H

			err = json.Unmarshal([]byte(unparseddata), &Testcase)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			data.Maxscore = len(Testcase)

			table.Structure[data.Id] = data
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": table,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type AssignedData struct {
	Id       string
	Name     string
	Info     string
	Link     string
	Testcase []gin.H
	Pass     any
	Template bool
}

func GetAssignedData(ctx *gin.Context) {
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

	if ok {
		result, err := database.Connection.Query(
			`SELECT asm.*,js.pass  FROM account asw CROSS JOIN JSON_TABLE(asw.work, '$.*'

			COLUMNS (
			 id VARCHAR(50) PATH '$.id',
			 pass TINYINT(1) PATH '$.pass'
			)
		 
		 	) AS js INNER JOIN assignment asm ON asm.Id = js.id WHERE asw.Username = ?;`,
			claims["username"],
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		var assignments []AssignedData = []AssignedData{}

		for result.Next() {
			var assignment AssignedData
			var unparseddata string
			err := result.Scan(&assignment.Id, &assignment.Name, &assignment.Info, &assignment.Link, &unparseddata, &assignment.Template, &assignment.Pass)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}
			err = json.Unmarshal([]byte(unparseddata), &assignment.Testcase)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			assignments = append(assignments, assignment)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": assignments,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type Group struct {
	Groupname  string
	Assignment map[string]gin.H
	Active     bool
}

func GetGroupList(ctx *gin.Context) {
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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}
		result, err := database.Connection.Query(
			"SELECT * FROM grouplist",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		var groups []Group

		for result.Next() {
			var group Group
			var unparseddata string
			err := result.Scan(&group.Groupname, &unparseddata, &group.Active)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			err = json.Unmarshal([]byte(unparseddata), &group.Assignment)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			groups = append(groups, group)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": groups,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}

}

type Assignment struct {
	Id       string
	Name     string
	Info     string
	Link     string
	Testcase []gin.H
	Template bool
}

func GetAssignmentList(ctx *gin.Context) {
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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}
		result, err := database.Connection.Query(
			"SELECT * FROM assignment",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		assignments := map[string]Assignment{}

		for result.Next() {
			var assignment Assignment
			var unparseddata string
			err := result.Scan(&assignment.Id, &assignment.Name, &assignment.Info, &assignment.Link, &unparseddata, &assignment.Template)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			err = json.Unmarshal([]byte(unparseddata), &assignment.Testcase)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			assignments[assignment.Id] = assignment
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": assignments,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type UserData struct {
	Username  string
	Firstname string
	Lastname  string
	Groups    []string
}

func GetUserData(ctx *gin.Context) {
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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}
		result, err := database.Connection.Query(
			"SELECT Username,Firstname,Lastname,groups FROM account",
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		var Users []UserData

		for result.Next() {
			var User UserData
			var unparseddata string

			err := result.Scan(&User.Username, &User.Firstname, &User.Lastname, &unparseddata)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			err = json.Unmarshal([]byte(unparseddata), &User.Groups)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}
			Users = append(Users, User)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": Users,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type GroupAssignment struct {
	Id       string
	Name     string
	Info     string
	Link     string
	Testcase []gin.H
	Template bool
}

type GroupWork struct {
	Groupname string
	Assigned  []GroupAssignment
}

func GetGroupWork(ctx *gin.Context) {
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

	if ok {
		username := claims["username"]

		result, err := database.Connection.Query(`
		SELECT groupname,Id,Name,Info,Link,Testcase,Template
		FROM account
		JOIN grouplist ON JSON_CONTAINS(account.groups, JSON_QUOTE(grouplist.groupname))
		JOIN JSON_TABLE(
		  JSON_KEYS(grouplist.assignment),
		  "$[*]" COLUMNS(key_value varchar(255) PATH "$")
		) AS keys_table ON 1=1
		JOIN assignment ON keys_table.key_value = assignment.Id
		WHERE Username = ? AND grouplist.active = 1
		;`,
			username,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalServerError",
				"response": "error while fetching data",
			})
			return
		}

		groupworks := map[string]GroupWork{}

		for result.Next() {
			var unparseddata string
			var groupname string
			var groupassignment GroupAssignment
			err = result.Scan(&groupname, &groupassignment.Id, &groupassignment.Name, &groupassignment.Info, &groupassignment.Link, &unparseddata, &groupassignment.Template)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			err = json.Unmarshal([]byte(unparseddata), &groupassignment.Testcase)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			entry, ok := groupworks[groupname]
			if !ok {
				entry = GroupWork{
					Groupname: groupname,
				}
			}

			entry.Assigned = append(entry.Assigned, groupassignment)
			groupworks[groupname] = entry

		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": groupworks,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

func GetUserCode(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	Sol := ctx.Query("sol")
	Target := ctx.Query("target")

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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		sourceFileDirectory := constant.Directory + "/storage/branch/" + Target + "/in/"
		if _, err := os.Stat(sourceFileDirectory); os.IsNotExist(err) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Not found this user folder:" + err.Error(),
			})
			return
		}

		sourceByte, err := os.ReadFile(sourceFileDirectory + Sol + ".cpp")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while read source file" + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": string(sourceByte),
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

func GetTemplateCode(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	Sol := ctx.Query("sol")

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

	_, ok := token.Claims.(jwt.MapClaims)

	if ok {
		sourceFileDirectory := constant.Directory + "/storage/sol/" + Sol + "/template/"
		if _, err := os.Stat(sourceFileDirectory); os.IsNotExist(err) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Not found template folder:" + err.Error(),
			})
			return
		}

		sourceByte, err := os.ReadFile(sourceFileDirectory + Sol + ".tp")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while read source file" + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": string(sourceByte),
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type ResponseSelfData struct {
	Firstname string
	Lastname  string
	Groups    []string
	Work      gin.H
}

func GetSelfData(ctx *gin.Context) {
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

	if ok {

		result := database.Connection.QueryRow("SELECT Firstname,Lastname,groups,work FROM account WHERE Username = ?", claims["username"])

		err = result.Err()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "InternalError",
				"output": "Error while getting self data " + " \n " + err.Error(),
			})
			return
		}

		var selfdata ResponseSelfData

		var unparseddata string
		var unparseddata2 string

		err = result.Scan(&selfdata.Firstname, &selfdata.Lastname, &unparseddata, &unparseddata2)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "InternalError",
				"output": "Error while scanning self row data for" + " \n " + err.Error(),
			})
			return
		}

		err = json.Unmarshal([]byte(unparseddata), &selfdata.Groups)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "InternalError",
				"output": "Error while parsing self group data" + " \n " + err.Error(),
			})
			return
		}

		err = json.Unmarshal([]byte(unparseddata2), &selfdata.Work)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "InternalError",
				"output": "Error while parsing self work data" + " \n " + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": selfdata,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}

type ShortAssignment struct {
	Id       string
	Name     string
	Maxscore int
}

type UserInfo struct {
	Username  string
	Firstname string
	Lastname  string
	Score     int
	Work      map[string]Work
}

type Groupdata struct {
	Groupname  string
	Member     map[string]UserInfo
	Assignment map[string]ShortAssignment
}

func GetGroupData(ctx *gin.Context) {
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

	if ok {
		isadmin := claims["isadmin"]
		if isadmin == false || isadmin == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":   "Forbidden",
				"response": "You not allow to use this route",
			})
			return
		}

		result, err := database.Connection.Query(`
			SELECT groupname,Id,Name,Testcase
			FROM grouplist JOIN JSON_TABLE(
					JSON_KEYS(grouplist.assignment),
					"$[*]" COLUMNS(key_value varchar(255) PATH "$")
					) AS keys_table ON 1=1
			JOIN assignment ON keys_table.key_value = assignment.Id;
		`)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		GroupDatas := map[string]Groupdata{}

		for result.Next() {
			var assignment ShortAssignment
			var groupname string
			var unparseddata string

			err := result.Scan(&groupname, &assignment.Id, &assignment.Name, &unparseddata)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while processing data",
				})
				fmt.Println(err.Error())
				return
			}

			var Testcase []gin.H

			err = json.Unmarshal([]byte(unparseddata), &Testcase)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status": "InternalError",
					"output": "Error while parsing data" + " \n " + err.Error(),
				})
				return
			}

			assignment.Maxscore = len(Testcase)

			entry, ok := GroupDatas[groupname]
			if !ok {
				entry = Groupdata{
					Groupname:  groupname,
					Member:     make(map[string]UserInfo),
					Assignment: map[string]ShortAssignment{},
				}
			}

			entry.Assignment[assignment.Id] = assignment
			GroupDatas[groupname] = entry
		}

		result, err = database.Connection.Query(`
		SELECT groupname,Username,Firstname,Lastname,work
		FROM account
		JOIN grouplist ON JSON_CONTAINS(account.groups, JSON_QUOTE(grouplist.groupname))
		JOIN JSON_TABLE(
		  JSON_KEYS(grouplist.assignment),
		  "$[*]" COLUMNS(key_value varchar(255) PATH "$")
		) AS keys_table ON 1=1
		JOIN assignment ON keys_table.key_value = assignment.Id
		;`)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":   "InternalError",
				"response": "Error while fetching data",
			})
			return
		}

		for result.Next() {
			var user UserInfo
			var groupname string
			var unparseddata string
			result.Scan(&groupname, &user.Username, &user.Firstname, &user.Lastname, &unparseddata)

			err = json.Unmarshal([]byte(unparseddata), &user.Work)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Error while parsing data",
				})
				fmt.Println(err.Error())
				return
			}

			entry, ok := GroupDatas[groupname]
			if !ok {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":   "InternalError",
					"response": "Invalid Data Group does not exist",
				})
				return
			}

			_, ok = entry.Member[user.Username]

			if !ok {

				user.Score = 0

				for _, v := range user.Work {
					if v.Pass {
						user.Score++
					}
				}
				entry.Member[user.Username] = user

			}

		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": GroupDatas,
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}
}
