package main

import (
	"fitness-tracker-api/testbackend/controllers"
	"fitness-tracker-api/testbackend/database"
	"log"
	"net/http"
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

	liftController := controllers.NewLiftController(db)
	cardioController := controllers.NewCardioController(db)
	workoutController := controllers.NewWorkoutController(db)
	liftingLogController := controllers.NewLiftingLogController(db)
	cardioLogController := controllers.NewCardioLogController(db)

	lift_group := server.Group("/lifts")
	{
		lift_group.GET("/:id", liftController.GetLift)
		lift_group.GET("", liftController.GetAllLifts)
		lift_group.GET("/search", liftController.SearchLiftsByName)
		lift_group.DELETE("/:id", liftController.DeleteLift)
		lift_group.POST("", liftController.AddLift)
	}

	cardio_log_group := server.Group("/cardiolog")
	{
		cardio_log_group.GET("", cardioLogController.GetAllCardioLogs)
		cardio_log_group.GET("/:id", cardioLogController.GetCardioLogById)
		cardio_log_group.GET("/workout/:workoutid", cardioLogController.CardioLogsByWorkout)
		cardio_log_group.DELETE("/:id", cardioLogController.DeleteCardioLog)
		cardio_log_group.POST("", cardioLogController.AddCardioLog)
	}

	cardio_group := server.Group("/cardio")
	{
		cardio_group.GET("/:id", cardioController.GetCardioById)
		cardio_group.GET("", cardioController.GetAllCardio)
		cardio_group.GET("/search", cardioController.SearchCardioByName)
		cardio_group.DELETE("/:id", cardioController.DeleteCardio)
		cardio_group.POST("", cardioController.AddCardio)
	}

	workout_log_group := server.Group("/workoutlog")
	{
		workout_log_group.GET("", liftingLogController.GetAllWorkoutLogs)
		workout_log_group.GET("/:id", liftingLogController.GetWorkoutLogById)
		workout_log_group.GET("/workout/:workoutid", liftingLogController.LiftingLogsByWorkout)
		workout_log_group.DELETE("/:id", liftingLogController.DeleteWorkoutLog)
		workout_log_group.POST("", liftingLogController.AddWorkoutLog)
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
			ctx.AbortWithStatus(http.StatusInternalServerError)
			log.Print(err)
			return
		}
		err = tmp.Execute(ctx.Writer, nil)
		// ctx.HTML(http.StatusOK, "documentation.tmpl", nil)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			log.Print(err)
			return
		}
	})

	server.Run(":8080")

}
