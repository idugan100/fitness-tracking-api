package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type CardioLogController struct {
	DB *sql.DB
}

func NewCardioLogController(DB *sql.DB) *CardioLogController {
	return &CardioLogController{DB}
}

func (c *CardioLogController) GetAllCardioLogs(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioLogController) GetCardioLogById(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioLogController) CardioLogsByWorkout(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioLogController) DeleteCardioLog(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}

func (c *CardioLogController) AddCardioLog(ctx *gin.Context) {
	ctx.JSON(200, "hi")
}
