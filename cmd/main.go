package main

import (
	"fmt"
	"github.com/woozie-10/students-rest-api/config"
	"github.com/woozie-10/students-rest-api/database"
	_ "github.com/woozie-10/students-rest-api/docs"
	"github.com/woozie-10/students-rest-api/router"
	"log"
)

// @title Students REST API
// @version 1.0
// @description This is a simple REST API project that manages student information in a MongoDB database using the Go programming language and the Gin web framework.
// @host localhost:8081
// @BasePath /
// @schemes http

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	router.InitRouter()

	if err := router.Router.Run(fmt.Sprintf("0.0.0.0:%s", config.Config.GetString("server.port"))); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
