package controllers

import (
	"fitness-tracker-api/testbackend/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllWorkouts(ctx *gin.Context) {

	rows, err := db_connection.Query("SELECT * FROM Workouts")

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer rows.Close()

	var workouts []models.Workout

	for rows.Next() {
		var workout models.Workout
		err = rows.Scan(&workout.Id, &workout.Location, &workout.Notes, &workout.Date)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(500)
			return
		} else {
			workouts = append(workouts, workout)
		}
	}

	ctx.JSON(200, workouts)
}
