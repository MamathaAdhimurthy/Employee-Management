package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	// r.Run()
	r := gin.Default()

	models.ConnectDatabase() // new
	r.GET("/employees", controllers.FindEmployees)
	r.POST("/employees", controllers.CreateEmployee)
	r.GET("/employees/:id", controllers.FindEmployee)
	r.PATCH("/employees/:id", controllers.UpdateEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)
	// r.GET("/books", controllers.FindBooks)
	// r.POST("/books", controllers.CreateBook)

	err := r.Run()
	if err != nil {
		return
	}
}
