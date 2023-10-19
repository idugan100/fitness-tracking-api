package controllers

import (
	"fitness-tracker-api/testbackend/models"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCardioById(ctx *gin.Context) {
	//validate id input
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	//get result from database and read into struct
	row, err := db_connection.Query("SELECT * FROM Cardio WHERE id=?", id)

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

func GetAllCardio(ctx *gin.Context) {
	rows, err := db_connection.Query("SELECT * FROM Cardio")

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

func SearchCardioByName(ctx *gin.Context) {
	search := ctx.Query("name")
	if search == "" {
		ctx.JSON(400, "missing name query parameter")
	}

	rows, err := db_connection.Query("SELECT * FROM Cardio WHERE Name LIKE ?", "%"+search+"%")
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

func DeleteCardio(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	_, err = db_connection.Exec("DELETE FROM Cardio WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(200, "cardio successfully deleted")
}

func AddCardio(ctx *gin.Context) {
	var cardio models.Cardio
	err := ctx.BindJSON(&cardio)
	if err != nil {
		log.Print(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = db_connection.Exec("INSERT INTO Cardio (name) VALUES (?)", cardio.Name)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSON(201, "cardio sucessfully created")
}
