package main

import (
	"fitness-tracker-api/testbackend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	controllers.Connect_to_database()
	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1"})

	lift_group := server.Group("/lifts")
	{
		lift_group.GET("/:id", controllers.GetLift)
		lift_group.GET("", controllers.GetAllLifts)
		lift_group.GET("/search", controllers.SearchLiftsByName)
		lift_group.DELETE("/delete/:id", controllers.DeleteLift)
		lift_group.POST("", controllers.AddLift)
	}

	cardio_group := server.Group("/cardio")
	{
		cardio_group.GET("/:id", controllers.GetCardioById)
		cardio_group.GET("", controllers.GetAllCardio)
		cardio_group.GET("/search", controllers.SearchCardioByName)
		cardio_group.DELETE("/:id", controllers.DeleteCardio)
		cardio_group.POST("", controllers.AddCardio)
	}

	workout_group := server.Group("/workouts")
	{
		workout_group.GET("", controllers.GetAllWorkouts)
		workout_group.GET("/:id", controllers.GetWorkout)
		workout_group.DELETE("/:id", controllers.DeleteWorkout)
		workout_group.POST("", controllers.AddWorkout)
	}

	server.Run(":8080")

}
