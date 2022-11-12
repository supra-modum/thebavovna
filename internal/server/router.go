package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	// in production
	//router.Use(static.Serve("/", static.LocalFile("./public/dist", true)))

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	api := router.Group("/api")

	{
		api.POST("/signin", signIn)
		api.POST("/signup", signUp)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
