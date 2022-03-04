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
		log.Fatal(err.Error())
		return
	}
	// open database
	dbEngine, err := common.OpenDB(config.DBName)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	// create gin-engine
	webEngine := gin.Default()

	// open services
	user.OpenService(dbEngine)
	thread.OpenService(dbEngine)
	route.OpenService(webEngine)
	templates.OpenService(webEngine)

	webEngine.Run(config.Adress)
}
