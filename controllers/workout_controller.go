package controllers

import (
	"fitness-tracker-api/testbackend/database"
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

	row, err := database.DB_connection.Query("SELECT * FROM Workouts WHERE id=?", id)

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

	rows, err := database.DB_connection.Query("SELECT * FROM Workouts")

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

func DeleteWorkout(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(400, "invalid id parameter")
		return
	}
	res, err := database.DB_connection.Exec("DELETE FROM Workouts WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	rows_deleted, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	if rows_deleted == 0 {
		log.Print("Unsuccessful Deletion - Resource Not Found")
		ctx.AbortWithStatus(410)
		return
	}

	ctx.JSON(200, "workout sucessfully deleted")
}

func AddWorkout(ctx *gin.Context) {
	var workout models.Workout
	err := ctx.BindJSON(&workout)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(410, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB_connection.Exec("INSERT INTO Workouts (Location, Notes) VALUES (?, ?)", &workout.Location, &workout.Notes)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(201, "workout sucessfully created")

}
