package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LiftingLogController struct {
	DB *sql.DB
}

func (lc *LiftingLogController) GetAllWorkoutLogs(ctx *gin.Context) {
	rows, err := lc.DB.Query("SELECT * FROM LiftingLog")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer rows.Close()

	var lifting_log_list []models.LiftingLog

	for rows.Next() {
		var lifting_log models.LiftingLog
		err = rows.Scan(&lifting_log.Id, &lifting_log.LiftId, &lifting_log.Weight, &lifting_log.Sets, &lifting_log.Reps, &lifting_log.WorkoutId)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(500)
			return
		}
		lifting_log_list = append(lifting_log_list, lifting_log)
	}
	ctx.JSON(200, lifting_log_list)
}

func (lc *LiftingLogController) GetWorkoutLogById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(400, "invalid id parameter")
		return
	}
	row, err := lc.DB.Query("SELECT * FROM LIFTINGLOG WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	var lifting_log models.LiftingLog
	is_found := row.Next()

	if !is_found {
		ctx.AbortWithStatus(404)
		return
	}

	err = row.Scan(&lifting_log.Id, &lifting_log.LiftId, &lifting_log.Weight, &lifting_log.Sets, &lifting_log.Reps, &lifting_log.WorkoutId)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(200, lifting_log)
}

func (lc *LiftingLogController) LiftingLogsByWorkout(ctx *gin.Context) {
	workout_id, err := strconv.Atoi(ctx.Param("workoutid"))
	log.Print(workout_id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(400, "invalid workout id")
		return
	}
	rows, err := lc.DB.Query("SELECT * FROM LiftingLog WHERE workout_id=?", workout_id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer rows.Close()

	var lifting_logs []models.LiftingLog

	for rows.Next() {
		log.Print("hi")
		var lifting_log models.LiftingLog
		err = rows.Scan(&lifting_log.Id, &lifting_log.LiftId, &lifting_log.Reps, &lifting_log.Sets, &lifting_log.Weight, &lifting_log.WorkoutId)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(500)
			return
		}
		lifting_logs = append(lifting_logs, lifting_log)
	}

	if len(lifting_logs) == 0 {
		ctx.AbortWithStatus(404)
		return
	}
	ctx.JSON(200, lifting_logs)

}

func (lc *LiftingLogController) DeleteWorkoutLog(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(400, "invalid id parameter")
		return
	}
	res, err := lc.DB.Exec("DELETE FROM LiftingLog WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	rows_affected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	if rows_affected == 0 {
		log.Print("unsuccessful deletion attempt")
		ctx.AbortWithStatus(410)
		return
	}
	ctx.JSON(200, "deletion successful")
}

func (lc *LiftingLogController) AddWorkoutLog(ctx *gin.Context) {
	var lifting_log models.LiftingLog
	err := ctx.BindJSON(&lifting_log)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	//check if lift id and workout id exists

	if lifting_log.Sets <= 0 {
		ctx.AbortWithStatusJSON(400, "sets must be positive")
		return
	}
	if lifting_log.Weight < 0 {
		ctx.AbortWithStatusJSON(400, "weight cannot be negative")
		return
	}
	if lifting_log.Reps <= 0 {
		ctx.AbortWithStatusJSON(400, "reps must be postive")
		return
	}

	res, err := lc.DB.Query("SELECT * FROM Workouts where id=?", lifting_log.WorkoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	if !res.Next() {
		ctx.AbortWithStatusJSON(400, "workout id does not exist")
		return
	}
	res.Close()

	res, err = lc.DB.Query("SELECT * FROM Lifts where id=?", lifting_log.LiftId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	if !res.Next() {
		ctx.AbortWithStatusJSON(400, "lift id does not exist")
		return
	}
	res.Close()

	_, err = lc.DB.Exec("INSERT INTO LiftingLog (lift_id, weight, sets, reps, workout_id) values (?,?,?,?,?)", lifting_log.LiftId, lifting_log.Weight, lifting_log.Sets, lifting_log.Reps, lifting_log.WorkoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	ctx.Status(201)
}
