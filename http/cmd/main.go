package main

import (
	"fmt"
	"focus/activity"
	"focus/impl"
	"log"

	"github.com/benbjohnson/clock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.LoadHTMLGlob("cmd/*.html")
	router.StaticFile("/main.js", "cmd/main.js")

	db, err := gorm.Open(sqlite.Open("activity/data/db.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        return
    }

	clock := clock.New()
	activityRepository := activity.NewSQLite(db, clock)
	if err != nil {
		log.Fatal(err)
	}
	focus := impl.New(activityRepository, clock)
	controller := &ginWrapper{focus: focus}

	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.DELETE("/", controller.Delete)
	router.GET("/health", controller.Health)

	port := 8080
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}