package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PORT string = "3000"
var ADDR string = "0.0.0.0:"+PORT 

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

	fmt.Println("Server is running on ",ADDR)
	app.Run(ADDR)
}
