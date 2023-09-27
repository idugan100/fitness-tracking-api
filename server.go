package main

import (
	"fitness-tracker-api/testbackend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	controllers.Connect_to_database()
	server := gin.Default()

	lift_group := server.Group("/lifts")
	{
		lift_group.GET("/:id", controllers.GetLift)
		lift_group.GET("", controllers.GetAllLifts)
		lift_group.GET("/search", controllers.SearchLiftsByName)
		lift_group.DELETE("/delete/:id", controllers.DeleteLift)
	}

	server.Run(":8080")

}
