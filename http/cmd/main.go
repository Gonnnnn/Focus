package main

import (
	"fmt"
	"focus/activity"
	"focus/impl"
	"log"

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

	activityRepository := activity.NewSQLite(db)
	if err != nil {
		log.Fatal(err)
	}
	focus := impl.New(activityRepository)
	controller := &ginWrapper{focus: focus}

	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.GET("/health", controller.Health)
	// router.GET("/terminate", terminator.terminate)

	port := 8080
	log.Printf("Listening on port %d", port)
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}