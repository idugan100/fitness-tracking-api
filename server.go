package main

import (
	"fitness-tracker-api/testbackend/controllers"
	"fitness-tracker-api/testbackend/database"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Print("Error opening database")
		log.Print(err)
		return
	}
	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1"})
	// server.LoadHTMLGlob("./templates/*")

	liftController := &controllers.LiftController{DB: db}
	cardioController := &controllers.CardioController{DB: db}
	workoutController := &controllers.WorkoutController{DB: db}

	lift_group := server.Group("/lifts")
	{
		lift_group.GET("/:id", liftController.GetLift)
		lift_group.GET("", liftController.GetAllLifts)
		lift_group.GET("/search", liftController.SearchLiftsByName)
		lift_group.DELETE("/:id", liftController.DeleteLift)
		lift_group.POST("", liftController.AddLift)
	}

	cardio_group := server.Group("/cardio")
	{
		cardio_group.GET("/:id", cardioController.GetCardioById)
		cardio_group.GET("", cardioController.GetAllCardio)
		cardio_group.GET("/search", cardioController.SearchCardioByName)
		cardio_group.DELETE("/:id", cardioController.DeleteCardio)
		cardio_group.POST("", cardioController.AddCardio)
	}

	workout_group := server.Group("/workouts")
	{
		workout_group.GET("", workoutController.GetAllWorkouts)
		workout_group.GET("/:id", workoutController.GetWorkout)
		workout_group.DELETE("/:id", workoutController.DeleteWorkout)
		workout_group.POST("", workoutController.AddWorkout)
	}
	server.GET("/documentation", func(ctx *gin.Context) {
		tmp, err := template.ParseFiles("./templates/documentation.tmpl")
		if err != nil {
			ctx.AbortWithStatus(500)
			log.Print(err)
			return
		}
		err = tmp.Execute(ctx.Writer, nil)
		// ctx.HTML(http.StatusOK, "documentation.tmpl", nil)
		if err != nil {
			ctx.AbortWithStatus(500)
			log.Print(err)
			return
		}
	})

	server.Run(":8080")

}
