package controllers

import (
	"fitness-tracker-api/testbackend/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWorkout(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(500, "invalid id parameter")
		return
	}

	row, err := db_connection.Query("SELECT * FROM Workouts WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer row.Close()

	row_found := row.Next()
	if !row_found {
		ctx.AbortWithStatus(404)
		return
	}

	var workout models.Workout
	err = row.Scan(&workout.Id, &workout.Location, &workout.Notes, &workout.Date)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
	}

	ctx.JSON(200, workout)
}

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
