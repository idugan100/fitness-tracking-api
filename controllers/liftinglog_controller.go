package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type LiftingLogController struct {
	DB *sql.DB
}

func (*LiftingLogController) GetAllWorkoutLogs(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*LiftingLogController) GetWorkoutLogById(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*LiftingLogController) WorkoutLogsByWorkout(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*LiftingLogController) DeleteWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

func (*LiftingLogController) AddWorkoutLog(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}
