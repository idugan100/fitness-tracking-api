package main

import (
	"fitness-tracker-api/testbackend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	controllers.Connect_to_database()
	server := gin.Default()

	server.GET("lift/:id", controllers.GetLift)
	server.GET("lifts", controllers.GetAllLifts)
	server.GET("lifts/search", controllers.SearchLiftsByName)
	server.Run(":8080")

}
