package main

import (
	"fmt"
	"focus/activity"
	"focus/impl"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lotusdblabs/lotusdb/v2"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	options := lotusdb.DefaultOptions
	options.DirPath = "/tmp/lotusdb_basic"
	db, err := lotusdb.Open(options)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	activityRepository := activity.NewLotus()
	focus := impl.New(activityRepository)
	controller := &ginWrapper{focus: focus}

	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.GET("/health", controller.Health)

	port := 8080
	log.Printf("Listening on port %d", port)
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	router.Run(fmt.Sprintf(":%d", port))
}
