package main

import (
	"fmt"
	"grader/auth"
	"grader/compiler"
	"grader/database"
	"grader/fetchdata"
	"grader/postdata"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	err := database.Init("root@/graderdb")

	if err != nil {
		fmt.Println("Error while connect to database server due to :" + err.Error())
		return
	}

	fmt.Println("Connected to database server")

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	app.Use(cors.New(config))

	app.SetTrustedProxies(nil)

	app.Static("/assets", "./web/serve/assets")
	app.NoRoute(func(ctx *gin.Context) {
		ctx.File("./web/serve/index.html")
	})

	app.POST("/api/sentry/updatetemplate", postdata.UpdateTemplate)
	app.POST("/api/sentry/updateassignment", postdata.UpdateAssignment)
	app.POST("/api/sentry/createassignment", postdata.CreateAssignment)
	app.POST("/api/sentry/updateusergroup", postdata.UpdateUserGroup)
	app.POST("/api/sentry/updatetestcase", postdata.UpdateTestcase)
	app.POST("/api/sentry/updategroupassignment", postdata.UpdateGroupAssignment)
	app.POST("/api/sentry/updategroupstate", postdata.UpdateGroupState)
	app.POST("/api/sentry/removegroup", postdata.RemoveGroup)
	app.POST("/api/sentry/regist", postdata.CreateUser)
	app.POST("/api/sentry/creategroup", postdata.CreateGroup)
	app.POST("/api/compile", compiler.CompileRequest)
	app.POST("/api/authenticate", auth.Authenticate)
	app.GET("/api/accessgate", auth.AccessGate)
	app.GET("/api/getassignment", fetchdata.GetAssignedData)
	app.GET("/api/getgroupwork", fetchdata.GetGroupWork)
	app.GET("/api/sentry/getusertable", fetchdata.GetUserTable)
	app.GET("/api/sentry/getassignmentlist", fetchdata.GetAssignmentList)
	app.GET("/api/sentry/getgrouplist", fetchdata.GetGroupList)
	app.GET("/api/sentry/getuser", fetchdata.GetUserData)
	app.GET("/api/sentry/getusercode", fetchdata.GetUserCode)
	app.GET("/api/gettemplatecode", fetchdata.GetTemplateCode)
	app.GET("/api/sentry/getgroupdata", fetchdata.GetGroupData)
	app.GET("/api/getselfdata", fetchdata.GetSelfData)

	err = app.Run(":3000")

	if err != nil {
		fmt.Println("Error while starting server due to :" + err.Error())
	}
}
