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
