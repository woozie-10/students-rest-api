package main

import (
	"fmt"
	"github.com/woozie-10/students-rest-api/config"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/router"
	"log"
)

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
