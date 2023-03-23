package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates router object
	r := gin.Default()

	// Assign handler function to be called and return a string
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	// Create a group of routes behind /api
	api := r.Group("/api")

	// Return a JSON response
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Serve Static files
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Tell the webserver to listen on port 3000
	r.Run()
}
