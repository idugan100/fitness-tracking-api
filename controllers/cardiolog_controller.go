package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardioLogController struct {
	DB *sql.DB
}

func NewCardioLogController(DB *sql.DB) *CardioLogController {
	return &CardioLogController{DB}
}

func (c *CardioLogController) GetAllCardioLogs(ctx *gin.Context) {
	rows, err := c.DB.Query("SELECT * FROM CardioLog")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var cardioLogList []models.CardioLog
	for rows.Next() {
		var cardioLog models.CardioLog
		err := rows.Scan(&cardioLog.Id, &cardioLog.CardioId, &cardioLog.Time, &cardioLog.Distance, &cardioLog.WorkoutId)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		cardioLogList = append(cardioLogList, cardioLog)
	}

	ctx.JSON(http.StatusOK, cardioLogList)
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
