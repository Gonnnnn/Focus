package main

import (
	"fmt"
	"focus/impl"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	focus := impl.New()
	controller := &ginWrapper{focus: focus}

	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.GET("/health", controller.Health)

	port := 8080
	log.Printf("Listening on port %d", port)
	router.Run(fmt.Sprintf(":%d", port))
}
