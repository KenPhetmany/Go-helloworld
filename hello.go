package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates router object
	r := gin.Default()

	// Assign handler function to be called and return a string
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	// Tell the webserver to listen on port 3000
	r.Run(":3000")
}
