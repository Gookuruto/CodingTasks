package main

import (
	controller "github.com/Gookuruto/CodingTasks/Controller"
	"github.com/gin-gonic/gin"
)

//NewRouter setup gin router basic version if there would be more controllers each one should have separate startup functions(files)
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORS())

	routerController := new(controller.RouteController)

	router.GET("/routes", routerController.FindRoutes)

	return router
}

//CORS function to allow all origins in case there would be connected some frontend without proxy only for dev purposes use this default implementation
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
