package main

import (
	"chatboard/common"
	"chatboard/route"
	"chatboard/templates"
	"chatboard/thread"
	"chatboard/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	config, err := common.LoadConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// open logger
	var logMode common.LogOutputMode
	if config.LogToFile {
		logMode = common.LogOutputFile
	} else {
		logMode = common.LogOutputScreen
	}
	err = common.SetupLogger(logMode)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// open database
	dbEngine, err := common.OpenDB(config.DBName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// create gin-engine
	webEngine := gin.Default()

	// open services
	user.OpenService(dbEngine)
	thread.OpenService(dbEngine)
	route.OpenService(webEngine)
	templates.OpenService(webEngine)

	common.LogInfo().Println("info test.")
	common.LogWarning().Println("warning test")
	common.LogError().Println("error test")

	webEngine.Run(config.Adress)
}
