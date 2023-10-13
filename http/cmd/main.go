package main

import (
	"fmt"
	"focus/activity"
	"focus/impl"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lotusdblabs/lotusdb/v2"
)

type terminator struct {
	sigChan chan os.Signal
}

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

	sigChan := make(chan os.Signal)
	go func(db *lotusdb.DB) {
		sig := <-sigChan
		if sig == os.Interrupt || sig == os.Kill {
			err := db.Close()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}(db)

	terminator := &terminator{sigChan: sigChan}
	activityRepository := activity.NewLotus()
	focus := impl.New(activityRepository)
	controller := &ginWrapper{focus: focus}

	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.GET("/health", controller.Health)
	router.GET("/terminate", terminator.terminate)

	port := 8080
	log.Printf("Listening on port %d", port)
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}

func (t *terminator) terminate(c *gin.Context) {
	t.sigChan <- os.Interrupt
}
