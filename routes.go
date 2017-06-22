package main

import (
	"github.com/gin-gonic/gin"
	todosHandler "github.com/justcoders/go-api-example/handlers/todos"
)

func registerRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) { c.String(200, "OK") })
	router.GET("/healthcheck", func(c *gin.Context) { c.String(200, "OK") })

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": Version,
			})
		})

		todos := api.Group("/todos")
		{
			todos.POST("/", todosHandler.Create)
			todos.GET("/", todosHandler.GetList)
			todos.GET("/:id", todosHandler.GetOne)
			todos.DELETE("/:id", todosHandler.GetOne)
		}
	}

	router.NoRoute()

	return
}
