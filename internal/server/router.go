package server

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public/dist", true)))

	// Create API route group
	api := router.Group("/api")

	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}

func Start() {
	router := setRouter()

	// Start listening and serving requests
	router.Run(":5000")
}
