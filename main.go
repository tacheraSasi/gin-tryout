package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PORT string = "3000"

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "Hello from gin",
			},
		)
	})

	fmt.Println("Server is running on port",PORT)
	app.Run(PORT)
}
