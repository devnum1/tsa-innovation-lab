package main

import (
	"tsa-innovation-lab/controllers"
	"tsa-innovation-lab/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()
	router.POST("/contacts", controllers.CreateContact)

	router.Run(":8080")
}
