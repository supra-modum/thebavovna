package server

import (
	"bavovna/internal/controllers"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	// in production
	//router.Use(static.Serve("/", static.LocalFile("./public/dist", true)))

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	api := router.Group("/api")

	api.POST("/register", controllers.Register)

	return router
}
