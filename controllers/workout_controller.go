package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkoutController struct {
	DB *sql.DB
}

func NewWorkoutController(DB *sql.DB) WorkoutController {
	return WorkoutController{DB}
}

func (w WorkoutController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "invalid id parameter")
		return
	}

	row, err := w.DB.Query("SELECT * FROM Workouts WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer row.Close()

	row_found := row.Next()
	if !row_found {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var workout models.Workout
	err = row.Scan(&workout.Id, &workout.Location, &workout.Notes, &workout.Date)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, workout)
}

func (w WorkoutController) GetAll(ctx *gin.Context) {

	rows, err := w.DB.Query("SELECT * FROM Workouts")

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var workouts []models.Workout
	var workout models.Workout
	for rows.Next() {
		err = rows.Scan(&workout.Id, &workout.Location, &workout.Notes, &workout.Date)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			workouts = append(workouts, workout)
		}
	}

	ctx.JSON(http.StatusOK, workouts)
}

func (w WorkoutController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id parameter")
		return
	}
	res, err := w.DB.Exec("DELETE FROM Workouts WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows_deleted, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if rows_deleted == 0 {
		log.Print("Unsuccessful Deletion - Resource Not Found")
		ctx.AbortWithStatus(http.StatusGone)
		return
	}

	_, err = w.DB.Exec("DELETE FROM CardioLog where workout_id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_, err = w.DB.Exec("DELETE FROM LiftingLog where workout_id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "workout sucessfully deleted")
}

func (w WorkoutController) Create(ctx *gin.Context) {
	var workout models.Workout
	err := ctx.BindJSON(&workout)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = w.DB.Exec("INSERT INTO Workouts (Location, Notes) VALUES (?, ?)", &workout.Location, &workout.Notes)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, "workout sucessfully created")

}
