package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type CardiologController struct {
	DB *sql.DB
}

func (c *CardioController) GetAllCardioLogs(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioController) GetCardioLogById(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioController) CardioLogsByWorkout(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioController) DeleteCardioLog(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioController) AddCardioLog(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}
