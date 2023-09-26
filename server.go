package main

import (
	"fitness-tracker-api/testbackend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	server := gin.Default()

	server.GET("lift/:id", controllers.GetLift)
	server.Run(":8080")

}
