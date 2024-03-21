package main

import (
	"fitness-tracker-api/testbackend/controllers"
	"fitness-tracker-api/testbackend/database"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
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

	var liftController controllers.ExerciseController = controllers.NewLiftController(db)
	var cardioController controllers.ExerciseController = controllers.NewCardioController(db)
	var workoutController controllers.ExerciseEventController = controllers.NewWorkoutController(db)
	var liftingLogController controllers.ExerciseLogController = controllers.NewLiftingLogController(db)
	var cardioLogController controllers.ExerciseLogController = controllers.NewCardioLogController(db)

	lift_group := server.Group("/lifts")
	{
		lift_group.GET("/:id", liftController.GetById)
		lift_group.GET("", liftController.GetAll)
		lift_group.GET("/search", liftController.Search)
		lift_group.DELETE("/:id", liftController.Delete)
		lift_group.POST("", liftController.Create)
	}

	cardio_log_group := server.Group("/cardiolog")
	{
		cardio_log_group.GET("", cardioLogController.GetAll)
		cardio_log_group.GET("/:id", cardioLogController.GetById)
		cardio_log_group.GET("/workout/:workoutid", cardioLogController.ByWorkout)
		cardio_log_group.DELETE("/:id", cardioLogController.Delete)
		cardio_log_group.POST("", cardioLogController.Create)
	}

	cardio_group := server.Group("/cardio")
	{
		cardio_group.GET("/:id", cardioController.GetById)
		cardio_group.GET("", cardioController.GetAll)
		cardio_group.GET("/search", cardioController.Search)
		cardio_group.DELETE("/:id", cardioController.Delete)
		cardio_group.POST("", cardioController.Create)
	}

	workout_log_group := server.Group("/workoutlog")
	{
		workout_log_group.GET("", liftingLogController.GetAll)
		workout_log_group.GET("/:id", liftingLogController.GetById)
		workout_log_group.GET("/workout/:workoutid", liftingLogController.ByWorkout)
		workout_log_group.DELETE("/:id", liftingLogController.Delete)
		workout_log_group.POST("", liftingLogController.Create)
	}

	workout_group := server.Group("/workouts")
	{
		workout_group.GET("", workoutController.GetAll)
		workout_group.GET("/:id", workoutController.GetById)
		workout_group.DELETE("/:id", workoutController.Delete)
		workout_group.POST("", workoutController.Create)
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
