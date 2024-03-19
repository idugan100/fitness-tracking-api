package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"net/http"
	"strconv"

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
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rows, err := c.DB.Query("SELECT * FROM CardioLog WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var cardioLog models.CardioLog
	err = rows.Scan(&cardioLog.Id, &cardioLog.CardioId, &cardioLog.Time, &cardioLog.Distance, &cardioLog.WorkoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cardioLog)
}

func (c *CardioLogController) CardioLogsByWorkout(ctx *gin.Context) {
	workoutIdString := ctx.Param("workoutid")
	workoutId, err := strconv.Atoi(workoutIdString)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rows, err := c.DB.Query("SELECT * FROM CardioLog WHERE workout_id=?", workoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
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

	if len(cardioLogList) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	ctx.JSON(http.StatusOK, cardioLogList)
}

func (c *CardioLogController) DeleteCardioLog(ctx *gin.Context) {
	idString := ctx.Param("id")
	cardioLogId, err := strconv.Atoi(idString)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := c.DB.Exec("DELETE FROM CardioLog WHERE id=?", cardioLogId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows_affected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if rows_affected == 0 {
		log.Print("unsuccessful deletion attempt")
		ctx.AbortWithStatus(http.StatusGone)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *CardioLogController) AddCardioLog(ctx *gin.Context) {
	var cardioLog models.CardioLog

	err := ctx.BindJSON(&cardioLog)
	if err != nil {
		log.Print(err)
		return
	}

	if cardioLog.Time <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "time must be positive")
		return
	} else if cardioLog.Distance <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "distance must be positive")
		return
	}

	res, err := c.DB.Query("SELECT * FROM Workouts where id=?", cardioLog.WorkoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !res.Next() {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "workout id does not exist")
		return
	}
	res.Close()

	res, err = c.DB.Query("SELECT * FROM Cardio where id=?", cardioLog.CardioId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !res.Next() {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "lift id does not exist")
		return
	}
	res.Close()

	_, err = c.DB.Exec("INSERT INTO CardioLog (cardio_id, time, distance, workout_id) VALUES (?,?,?,?)", cardioLog.CardioId, cardioLog.Time, cardioLog.Distance, cardioLog.WorkoutId)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
