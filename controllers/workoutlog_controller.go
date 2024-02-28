package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type WorkoutLogController struct {
	DB *sql.DB
}

func (*WorkoutLogController) GetAllWorkoutLogs(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*WorkoutLogController) GetWorkoutLogById(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*WorkoutLogController) WorkoutLogsByWorkout(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*WorkoutLogController) DeleteWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*WorkoutLogController) AddWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}
