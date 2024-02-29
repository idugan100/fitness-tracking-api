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
		ctx.AbortWithStatusJSON(500, "invalid id parameter")
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

func (lc *LiftingLogController) WorkoutLogsByWorkout(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (lc *LiftingLogController) DeleteWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (lc *LiftingLogController) AddWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}
