package compiler

import (
	"context"
	"encoding/json"
	"fmt"
	"grader/constant"
	"grader/database"
	"grader/postdata"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CompileRequest(ctx *gin.Context) {

	var compileData CompileRequestData

	err := ctx.BindJSON(&compileData) //แปลงข้อมูลจาก json ให้เป็น struct ของ go

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "DataError",
			"output": "compiler could not process data",
		})
		return
	}
	Runcode(ctx, &compileData)

}

type RequestGraderTestcase struct {
	Name  string `json:"name"`
	Stdin bool   `json:"stdin"`
}

type CompileRequestData struct {
	Id             string                  `json:"id"`
	Code           string                  `json:"code"`
	Ctype          string                  `json:"ctype"`
	Grader         bool                    `json:"grade"`
	GraderTestcase []RequestGraderTestcase `json:"gradeoption"`
	Integrated     bool                    `json:"integrated"`
	Sol            string                  `json:"solution"`
	Input          string                  `json:"input"`
	Islogin        bool                    `json:"islogin"`
	Username       string                  `json:"username"`
}

// dictinterprinter := map[string]bool{"python":true,"javascript":true}

func isGrader(compileData *CompileRequestData) bool {
	return compileData.Grader
}

func isSolutionEmpty(compileData *CompileRequestData) bool {
	return len(compileData.Sol) <= 0
}

func isTestcaseEmpty(compileData *CompileRequestData) bool {
	return len(compileData.GraderTestcase) <= 0
}

func ThrowException(ctx *gin.Context, statusCode int, messageDetail string) {

	var statusMessage string

	switch statusCode {
	case http.StatusBadRequest:
		statusMessage = "BadData"
	case http.StatusInternalServerError:
		statusMessage = "InternalError"
	case http.StatusRequestTimeout:
		statusMessage = "TimeoutError"
	default:
		statusMessage = "UnknownError"
	}

	ctx.JSON(statusCode, gin.H{
		"status": statusMessage,
		"output": messageDetail,
	})
}

func isGraderDataPassed(compileData *CompileRequestData) (bool, string) {
	if isSolutionEmpty(compileData) {
		return false, "Solution is empty"
	}
	if isTestcaseEmpty(compileData) {
		return false, "No Testcase data included"
	}
	return true, ""
}

var LangType map[string]map[string]string = map[string]map[string]string{
	"c++": {
		"file":     ".cpp",
		"type":     "compiler",
		"compiler": "g++",
		"argument": "-o",
	},
	"c": {
		"file":     ".c",
		"type":     "compiler",
		"compiler": "gcc",
		"argument": "-o",
	},
	"python": {
		"file":     ".py",
		"type":     "interpreter",
		"compiler": "python",
		"argument": "-u",
	},
	"javascript": {
		"file":     ".js",
		"type":     "interpreter",
		"compiler": "node",
		"argument": "",
	},
	"php": {
		"file":     ".php",
		"type":     "interpreter",
		"compiler": "php",
		"argument": "",
	},
	"lua": {
		"file":     ".lua",
		"type":     "interpreter",
		"compiler": "lua",
		"argument": "",
	},
}

func Runcode(ctx *gin.Context, compileData *CompileRequestData) { //change name function
	// var isinteprinter bool = false;

	// if(dictinterprinter[CompileRequestData.Ctype]){
	// 	isinteprinter=true;
	// }

	entry, ok := LangType[compileData.Ctype]

	if !ok {
		ThrowException(ctx, http.StatusBadRequest, "Unsupported language")
		return
	}

	if isGrader(compileData) {
		if passed, errMessage := isGraderDataPassed(compileData); passed {
			ThrowException(ctx, http.StatusBadRequest, errMessage)
			return
		}
	}
	var inPath string
	var outPath string
	if compileData.Islogin {
		loginDirectory := constant.Directory + "/storage/branch/" + compileData.Username
		inPath = loginDirectory + "/in/"
		outPath = loginDirectory + "/out/"
		if _, err := os.Stat(loginDirectory); os.IsNotExist(err) {
			err := postdata.CreateFolder(loginDirectory)
			if err != nil {
				ThrowException(ctx, http.StatusInternalServerError, "Error creating folder:"+err.Error())
				return
			}
		}
		if _, err := os.Stat(inPath); os.IsNotExist(err) {
			err := postdata.CreateFolder(inPath)
			if err != nil {
				ThrowException(ctx, http.StatusInternalServerError, "Error creating folder:"+err.Error())
				return
			}
		}
		if _, err := os.Stat(outPath); os.IsNotExist(err) {
			err := postdata.CreateFolder(outPath)
			if err != nil {
				ThrowException(ctx, http.StatusInternalServerError, "Error creating folder:"+err.Error())
				return
			}
		}
		if compileData.Grader {
			inPath = inPath + "code-" + compileData.Sol + entry["file"]
			outPath = outPath + "exe-" + compileData.Sol
		} else {
			inPath = inPath + "code-" + compileData.Id + entry["file"]
			outPath = outPath + "exe-" + compileData.Id
		}

	} else {
		inPath = constant.Directory + "/storage/tmp_run/in/code-" + compileData.Id + entry["file"]
		outPath = constant.Directory + "/storage/tmp_run/out/exe-" + compileData.Id
	}

	file, err := os.Create(inPath)
	if err != nil {
		ThrowException(ctx, http.StatusInternalServerError, "Failed to create temporary file")
		return
	}
	defer file.Close()

	_, err = file.WriteString(compileData.Code)
	if err != nil {
		ThrowException(ctx, http.StatusInternalServerError, "Failed to write C++ code to file")
		return
	}

	var cmd *exec.Cmd
	var output []byte

	if entry["type"] == "compiler" {
		cmd := exec.Command(entry["compiler"], entry["argument"], outPath+".exe", inPath) //เตรียมรัน
		output, err := cmd.CombinedOutput()                                               //รันจริง
		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Failed to compiled code:\n"+string(output)+" \n "+err.Error())
			return
		}
	}

	if entry["type"] == "interpreter" {
		if entry["argument"] != "" {
			cmd = exec.Command(entry["compiler"], entry["argument"], inPath)
		} else {
			cmd = exec.Command(entry["compiler"], inPath)
		}
	}

	if !compileData.Grader {
		if entry["type"] == "compiler" {
			cmd = exec.Command(outPath)
		}
		cmd.Stdin = strings.NewReader(compileData.Input)

		outChannel := make(chan string)
		errorChannel := make(chan error)

		go func() {
			output, err = cmd.CombinedOutput()
			if err != nil {
				errorChannel <- err
			} else {
				outChannel <- string(output)
			}
		}()

		timeout := 5 * time.Second

		ctxTimeout, cancel := context.WithTimeout(ctx, timeout) //return ออกมาเป็น channel
		defer cancel()

		select {
		case output := <-outChannel:
			if entry["type"] == "compiler" {
				err = os.Remove(outPath + ".exe")
			}
			var addtionError string = ""
			if err != nil {
				addtionError = err.Error()
			}

			ctx.JSON(http.StatusOK, gin.H{
				"status": "success",
				"output": addtionError + "\n" + string(output),
			})
			return

		case err := <-errorChannel:

			if err != nil {
				if entry["type"] == "compiler" {
					err2 := os.Remove(outPath + ".exe")
					if err2 != nil {
						ThrowException(ctx, http.StatusInternalServerError, "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n"+err.Error()+"\n and Failed to delete execute file"+err2.Error())
						return
					}
				}

				ThrowException(ctx, http.StatusInternalServerError, "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n"+err.Error())
				return

			}

		case <-ctxTimeout.Done():

			cmd.Process.Kill()
			var addtionError string = ""
			if entry["type"] == "compiler" {
				err = os.Remove(outPath + ".exe")
				if err != nil {
					addtionError = err.Error()
				}
			}

			ThrowException(ctx, http.StatusRequestTimeout, "Execution reach timeout (Infinite loop cause) "+addtionError)
			return
		}
	} else {
		row := database.Connection.QueryRow(`
		SELECT Id
		FROM account
		JOIN grouplist ON JSON_CONTAINS(account.groups, JSON_QUOTE(grouplist.groupname))
		JOIN JSON_TABLE(
		  JSON_KEYS(grouplist.assignment),
		  "$[*]" COLUMNS(key_value varchar(255) PATH "$")
		) AS keys_table ON 1=1
		JOIN assignment ON keys_table.key_value = assignment.Id
		WHERE Username = ? AND Id = ? AND grouplist.active = 1
		;`,
			compileData.Username,
			compileData.Sol,
		)
		var Id string

		if err := row.Scan(&Id); err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status": "Forbidden",
				"output": "Your work is no longer available",
			})
			fmt.Println(err.Error())
			return
		} else {
			row = nil
		}

		row = database.Connection.QueryRow(`
		SELECT COALESCE(workjs.attempt, 0) FROM account
		CROSS JOIN JSON_TABLE(
		  work,
		  "$.*" COLUMNS(attempt INT(4) PATH '$.attempt', Id varchar(255) PATH '$.Id')
		) AS workjs
		WHERE Username = ? AND workjs.Id = ?
		;`,
			compileData.Username,
			compileData.Sol,
		)

		var attempt int

		if err := row.Scan(&attempt); err != nil {
			attempt = 0
		}

		var results []gin.H
		directory := constant.Directory + "/storage/sol/"

		var updatedata gin.H

		var score int = 0

		for _, v := range compileData.GraderTestcase {

			cmd = exec.Command(outPath)
			if v.Stdin {
				stdin, err := os.ReadFile(directory + compileData.Sol + "/in/" + v.Name + ".in")
				if err != nil {
					ThrowException(ctx, http.StatusInternalServerError, v.Name+" Input file not found in storage")
					return
				}
				cmd.Stdin = strings.NewReader(string(stdin))
			}

			outChannel := make(chan string)
			errorChannel := make(chan error)

			go func() {
				output, err = cmd.CombinedOutput()
				if err != nil {
					errorChannel <- err
				} else {
					outChannel <- string(output)
				}
			}()

			timeout := 2 * time.Second

			ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			select {
			case output := <-outChannel:
				expected, pass, err := GraderCheck(output, directory+compileData.Sol+"/out/"+v.Name+".sol")

				if err != nil {
					ThrowException(ctx, http.StatusInternalServerError, v.Name+" Output file not found in storage")
					fmt.Println(err.Error())
					return
				}

				if pass {
					results = append(results, gin.H{
						"case":     v.Name,
						"status":   "pass",
						"expected": expected,
						"output":   string(output),
					})
					score++

				} else {
					results = append(results, gin.H{
						"case":     v.Name,
						"status":   "rejected",
						"expected": expected,
						"output":   string(output),
					})

				}

			case err := <-errorChannel:
				if err != nil {
					results = append(results, gin.H{
						"case":   v.Name,
						"status": "InternalError",
						"output": "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n" + string(output) + " \n " + err.Error(),
					})

				}
			case <-ctxTimeout.Done():

				cmd.Process.Kill()
				results = append(results, gin.H{
					"case":   v.Name,
					"status": "TimeError",
					"output": "Execution reach timeout (Infinite loop cause)",
				})
			}
		}

		if score >= len(compileData.GraderTestcase) {
			updatedata = gin.H{
				"Id":        compileData.Sol,
				"pass":      true,
				"timestamp": time.Now().Unix(),
				"score":     score,
				"attempt":   attempt + 1,
			}
		} else {
			updatedata = gin.H{
				"Id":        compileData.Sol,
				"pass":      false,
				"timestamp": time.Now().Unix(),
				"score":     score,
				"attempt":   attempt + 1,
			}
		}

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n"+err.Error())
			return
		}

		var unparseddata string

		result := database.Connection.QueryRow("SELECT work FROM account WHERE Username = ?",
			compileData.Username,
		)

		err = result.Err()

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error while getting data for update"+" \n "+err.Error())
			return
		}

		err = result.Scan(&unparseddata)

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error while scanning row data for update"+" \n "+err.Error())
			return
		}

		var parsedwork gin.H

		err = json.Unmarshal([]byte(unparseddata), &parsedwork)

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error while parsing data for update"+" \n "+err.Error())
			return
		}

		mergeddata := parsedwork

		mergeddata[compileData.Sol] = updatedata

		jsonbyte, err := json.Marshal(&mergeddata)

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error while stringfy data"+" \n "+err.Error())
			return
		}

		_, err = database.Connection.Exec("UPDATE account set work = ? WHERE Username = ?",
			string(jsonbyte),
			compileData.Username,
		)

		if err != nil {
			ThrowException(ctx, http.StatusInternalServerError, "Error while updating data"+" \n "+err.Error())
			return
		}

		if entry["type"] == "compiler" {
			err = os.Remove(outPath + ".exe")
			if err != nil {
				ThrowException(ctx, http.StatusInternalServerError, "Error while detaching execute file \n but still get your result : "+string(output)+" \n "+err.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"output": results,
			"work":   mergeddata,
		})
	}

}

func GraderCheck(in string, path string) (string, bool, error) {
	solvedContent, err := os.ReadFile(path)

	solvedString := string(solvedContent)

	if err != nil {
		return solvedString, false, err
	}

	if strings.TrimSpace(in) == strings.TrimSpace(solvedString) {
		return solvedString, true, nil
	} else {
		return solvedString, false, nil
	}
}

// func python(ctx *gin.Context, compileData *CompileRequestData) {

// 	if compileData.Grader && len(compileData.Sol) <= 0 {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"status": "BadData",
// 			"output": "Solution is empty",
// 		})
// 		return
// 	}

// 	if compileData.Grader && len(compileData.GraderTestcase) <= 0 {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"status": "BadData",
// 			"output": "No Testcase data included",
// 		})
// 		return
// 	}
// 	var inPath string
// 	var outPath string
// 	if compileData.Islogin {
// 		loginDirectory := constant.Directory + "/storage/branch/" + compileData.Username
// 		inPath = loginDirectory + "/in/"
// 		outPath = loginDirectory + "/out/"
// 		if _, err := os.Stat(loginDirectory); os.IsNotExist(err) {
// 			err := postdata.CreateFolder(loginDirectory)
// 			if err != nil {
// 				ctx.JSON(http.StatusInternalServerError, gin.H{
// 					"status":   "InternalError",
// 					"response": "Error creating folder:" + err.Error(),
// 				})
// 				return
// 			}
// 		}
// 		if _, err := os.Stat(inPath); os.IsNotExist(err) {
// 			err := postdata.CreateFolder(inPath)
// 			if err != nil {
// 				ctx.JSON(http.StatusInternalServerError, gin.H{
// 					"status":   "InternalError",
// 					"response": "Error creating folder:" + err.Error(),
// 				})
// 				return
// 			}
// 		}
// 		if _, err := os.Stat(outPath); os.IsNotExist(err) {
// 			err := postdata.CreateFolder(outPath)
// 			if err != nil {
// 				ctx.JSON(http.StatusInternalServerError, gin.H{
// 					"status":   "InternalError",
// 					"response": "Error creating folder:" + err.Error(),
// 				})
// 				return
// 			}
// 		}
// 		if compileData.Grader {
// 			inPath = inPath + "code-" + compileData.Sol + ".py"
// 		} else {
// 			inPath = inPath + "code-" + compileData.Id + ".py"
// 		}

// 	} else {
// 		inPath = constant.Directory + "/storage/tmp_run/in/code-" + compileData.Id + ".py"
// 	}

// 	file, err := os.Create(inPath)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": "InternalError",
// 			"output": "Failed to create temporary file",
// 		})
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(compileData.Code)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": "InternalError",
// 			"output": "Failed to write C++ code to file",
// 		})
// 		return
// 	}

// 	cmd := exec.Command("python", "-u", inPath) //เตรียมรัน

// 	if !compileData.Grader {

// 		cmd.Stdin = strings.NewReader(compileData.Input)

// 		outChannel := make(chan string)
// 		errorChannel := make(chan error)

// 		go func() {
// 			output, err := cmd.CombinedOutput()
// 			if err != nil {
// 				errorChannel <- err
// 			} else {
// 				outChannel <- string(output)
// 			}
// 		}()

// 		timeout := 5 * time.Second

// 		ctxTimeout, cancel := context.WithTimeout(ctx, timeout) //return ออกมาเป็น channel
// 		defer cancel()

// 		select {
// 		case output := <-outChannel:
// 			var addtionError string = ""
// 			ctx.JSON(http.StatusOK, gin.H{
// 				"status": "success",
// 				"output": addtionError + "\n" + string(output),
// 			})
// 			return

// 		case err := <-errorChannel:

// 			if err != nil {
// 				ctx.JSON(http.StatusInternalServerError, gin.H{
// 					"status": "InternalError",
// 					"output": "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n" + err.Error(),
// 				})
// 				return
// 			}

// 		case <-ctxTimeout.Done():

// 			cmd.Process.Kill()

// 			var addtionError string = ""
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "TimeError",
// 				"output": "Execution reach timeout (Infinite loop cause) " + addtionError,
// 			})
// 			return
// 		}
// 	} else {
// 		row := database.Connection.QueryRow(`
// 		SELECT Id
// 		FROM account
// 		JOIN grouplist ON JSON_CONTAINS(account.groups, JSON_QUOTE(grouplist.groupname))
// 		JOIN JSON_TABLE(
// 		  JSON_KEYS(grouplist.assignment),
// 		  "$[*]" COLUMNS(key_value varchar(255) PATH "$")
// 		) AS keys_table ON 1=1
// 		JOIN assignment ON keys_table.key_value = assignment.Id
// 		WHERE Username = ? AND Id = ? AND grouplist.active = 1
// 		;`,
// 			compileData.Username,
// 			compileData.Sol,
// 		)
// 		var Id string

// 		if err := row.Scan(&Id); err != nil {
// 			ctx.JSON(http.StatusForbidden, gin.H{
// 				"status": "Forbidden",
// 				"output": "Your work is no longer available",
// 			})
// 			fmt.Println(err.Error())
// 			return
// 		} else {
// 			row = nil
// 		}

// 		row = database.Connection.QueryRow(`
// 		SELECT COALESCE(workjs.attempt, 0) FROM account
// 		CROSS JOIN JSON_TABLE(
// 		  work,
// 		  "$.*" COLUMNS(attempt INT(4) PATH '$.attempt', Id varchar(255) PATH '$.Id')
// 		) AS workjs
// 		WHERE Username = ? AND workjs.Id = ?
// 		;`,
// 			compileData.Username,
// 			compileData.Sol,
// 		)

// 		var attempt int

// 		if err := row.Scan(&attempt); err != nil {
// 			attempt = 0
// 		}

// 		var results []gin.H
// 		directory := constant.Directory + "/storage/sol/"

// 		var updatedata gin.H

// 		var score int = 0

// 		for _, v := range compileData.GraderTestcase {

// 			cmd = exec.Command(outPath)
// 			if v.Stdin {
// 				stdin, err := os.ReadFile(directory + compileData.Sol + "/in/" + v.Name + ".in")
// 				if err != nil {
// 					ctx.JSON(http.StatusInternalServerError, gin.H{
// 						"status": "InternalError",
// 						"output": v.Name + " Input file not found in storage",
// 					})
// 					return
// 				}
// 				cmd.Stdin = strings.NewReader(string(stdin))
// 			}

// 			outChannel := make(chan string)
// 			errorChannel := make(chan error)

// 			var output string
// 			go func() {
// 				output, err := cmd.CombinedOutput()
// 				if err != nil {
// 					errorChannel <- err
// 				} else {
// 					outChannel <- string(output)
// 				}
// 			}()

// 			timeout := 2 * time.Second

// 			ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
// 			defer cancel()
// 			select {
// 			case output := <-outChannel:
// 				expected, pass, err := GraderCheck(output, directory+compileData.Sol+"/out/"+v.Name+".sol")

// 				if err != nil {
// 					ctx.JSON(http.StatusInternalServerError, gin.H{
// 						"status": "InternalError",
// 						"output": v.Name + " Output file not found in storage",
// 					})
// 					fmt.Println(err.Error())
// 					return
// 				}

// 				if pass {
// 					results = append(results, gin.H{
// 						"case":     v.Name,
// 						"status":   "pass",
// 						"expected": expected,
// 						"output":   string(output),
// 					})
// 					score++

// 				} else {
// 					results = append(results, gin.H{
// 						"case":     v.Name,
// 						"status":   "rejected",
// 						"expected": expected,
// 						"output":   string(output),
// 					})

// 				}

// 			case err := <-errorChannel:
// 				if err != nil {
// 					results = append(results, gin.H{
// 						"case":   v.Name,
// 						"status": "InternalError",
// 						"output": "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n" + string(output) + " \n " + err.Error(),
// 					})

// 				}
// 			case <-ctxTimeout.Done():

// 				cmd.Process.Kill()
// 				results = append(results, gin.H{
// 					"case":   v.Name,
// 					"status": "TimeError",
// 					"output": "Execution reach timeout (Infinite loop cause)",
// 				})
// 			}
// 		}

// 		if score >= len(compileData.GraderTestcase) {
// 			updatedata = gin.H{
// 				"Id":        compileData.Sol,
// 				"pass":      true,
// 				"timestamp": time.Now().Unix(),
// 				"score":     score,
// 				"attempt":   attempt + 1,
// 			}
// 		} else {
// 			updatedata = gin.H{
// 				"Id":        compileData.Sol,
// 				"pass":      false,
// 				"timestamp": time.Now().Unix(),
// 				"score":     score,
// 				"attempt":   attempt + 1,
// 			}
// 		}

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error \nprogram doesn't exit with 0 value\n or program terminate because try to access non allocate memory block\n" + err.Error(),
// 			})
// 			return
// 		}

// 		var unparseddata string

// 		result := database.Connection.QueryRow("SELECT work FROM account WHERE Username = ?",
// 			compileData.Username,
// 		)

// 		err = result.Err()

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error while getting data for update" + " \n " + err.Error(),
// 			})
// 			return
// 		}

// 		err = result.Scan(&unparseddata)

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error while scanning row data for update" + " \n " + err.Error(),
// 			})
// 			return
// 		}

// 		var parsedwork gin.H

// 		err = json.Unmarshal([]byte(unparseddata), &parsedwork)

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error while parsing data for update" + " \n " + err.Error(),
// 			})
// 			return
// 		}

// 		mergeddata := parsedwork

// 		mergeddata[compileData.Sol] = updatedata

// 		jsonbyte, err := json.Marshal(&mergeddata)

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error while stringfy data" + " \n " + err.Error(),
// 			})
// 			return
// 		}

// 		_, err = database.Connection.Exec("UPDATE account set work = ? WHERE Username = ?",
// 			string(jsonbyte),
// 			compileData.Username,
// 		)

// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"status": "InternalError",
// 				"output": "Error while updating data" + " \n " + err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"status": "success",
// 			"output": results,
// 			"work":   mergeddata,
// 		})
// 	}

// }
