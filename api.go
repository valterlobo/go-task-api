package main

import (
	"github.com/gin-gonic/gin"
	"lobo.tech/task/controller"
)

func main() {

	r := gin.Default()
	controller.NewTaskController(r)
	r.GET("/ping", func(r *gin.Context) {
		r.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
