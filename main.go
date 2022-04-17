//package fasthttp_api
package main

import (
	"fasthttp_api/controllers"
	"fasthttp_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.SetTrustedProxies([]string{"192.168.1.2"})

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
