package main

import (
	"chatboard/common"
	"chatboard/thread"
	"chatboard/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	dbEngine, err := common.OpenDB(config.DBName, config.DBSync)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	webEngine := gin.Default()
	user.OpenService(dbEngine)
	thread.OpenService(dbEngine)
	webEngine.Run(config.Adress)
}
