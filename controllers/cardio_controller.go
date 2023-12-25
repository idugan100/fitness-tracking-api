package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

type CardioController struct {
	DB *sql.DB
}

func (cc *CardioController) GetCardioById(ctx *gin.Context) {
	//validate id input
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(500, "invalid id parameter")
		return
	}

	//get result from database and read into struct
	row, err := cc.DB.Query("SELECT * FROM Cardio WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer row.Close()

	var cardio models.Cardio
	row_found := row.Next()

	if !row_found {
		ctx.AbortWithStatus(404)
		return
	}
	err = row.Scan(&cardio.Id, &cardio.Name)
	if err != nil {
		log.Print(row, err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(200, cardio)
}

func (cc *CardioController) GetAllCardio(ctx *gin.Context) {
	rows, err := cc.DB.Query("SELECT * FROM Cardio")

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer rows.Close()

	var cardio_list []models.Cardio

	for rows.Next() {
		var cardio models.Cardio
		err = rows.Scan(&cardio.Id, &cardio.Name)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(500)
			return
		} else {
			cardio_list = append(cardio_list, cardio)
		}
	}

	ctx.JSON(200, cardio_list)
}

func (cc *CardioController) SearchCardioByName(ctx *gin.Context) {
	search := ctx.Query("name")
	if search == "" {
		ctx.JSON(400, "missing name query parameter")
	}

	rows, err := cc.DB.Query("SELECT * FROM Cardio WHERE Name LIKE ?", "%"+search+"%")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer rows.Close()

	var cardio_list []models.Cardio
	for rows.Next() {
		var cardio models.Cardio
		err = rows.Scan(&cardio.Id, &cardio.Name)
		if err != nil {
			log.Print(cardio, err)
			ctx.AbortWithStatus(500)
			return
		}
		cardio_list = append(cardio_list, cardio)
	}
	ctx.JSON(200, cardio_list)
}

func (cc *CardioController) DeleteCardio(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	res, err := cc.DB.Exec("DELETE FROM Cardio WHERE id=?", id)
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
		log.Print("Unsuccessful Deletion - Resource Not Found")
		ctx.AbortWithStatus(410)
		return
	}

	ctx.JSON(200, "cardio successfully deleted")
}

func (cc *CardioController) AddCardio(ctx *gin.Context) {
	var cardio models.Cardio
	err := ctx.BindJSON(&cardio)
	if err != nil {
		log.Print(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = cc.DB.Exec("INSERT INTO Cardio (name) VALUES (?)", cardio.Name)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSON(201, "cardio sucessfully created")
}
