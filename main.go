package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/justcoders/go-api-example/db"
	"github.com/justcoders/go-api-example/middlewares"
)

const (
	Port    = "3000"
	Version = "0.0.1"
)

func init() {
	db.Connect("mongodb://localhost:27017/sample_db")
}

func main() {

	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORS)

	registerRoutes(router)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Start listening on " + port)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
