package main

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	// initialize gin
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Content-Type",
		MaxAge:         60 * time.Second,
	}))

	// Add our api routes
	LoadAPIRoutes(router)

	router.Run(":4000")
}
