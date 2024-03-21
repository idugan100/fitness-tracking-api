package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CardioController struct {
	DB *sql.DB
}

func NewCardioController(DB *sql.DB) CardioController {
	return CardioController{DB}
}
func (c CardioController) GetById(ctx *gin.Context) {
	//validate id input
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "invalid id parameter")
		return
	}

	//get result from database and read into struct
	row, err := c.DB.Query("SELECT * FROM Cardio WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var cardio models.Cardio
	row_found := row.Next()

	if !row_found {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = row.Scan(&cardio.Id, &cardio.Name)
	if err != nil {
		log.Print(row, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, cardio)
}

func (c CardioController) GetAll(ctx *gin.Context) {
	rows, err := c.DB.Query("SELECT * FROM Cardio")

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cardio_list []models.Cardio
	var cardio models.Cardio

	for rows.Next() {
		err = rows.Scan(&cardio.Id, &cardio.Name)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			cardio_list = append(cardio_list, cardio)
		}
	}

	ctx.JSON(http.StatusOK, cardio_list)
}

func (c CardioController) Search(ctx *gin.Context) {
	search := ctx.Query("name")
	if search == "" {
		ctx.JSON(http.StatusBadRequest, "missing name query parameter")
	}

	rows, err := c.DB.Query("SELECT * FROM Cardio WHERE Name LIKE ?", "%"+search+"%")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cardio models.Cardio
	var cardio_list []models.Cardio
	for rows.Next() {
		err = rows.Scan(&cardio.Id, &cardio.Name)
		if err != nil {
			log.Print(cardio, err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		cardio_list = append(cardio_list, cardio)
	}
	ctx.JSON(http.StatusOK, cardio_list)
}

func (c CardioController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	res, err := c.DB.Exec("DELETE FROM Cardio WHERE id=?", id)
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
		log.Print("Unsucessful Deletion - Resource Not Found")
		ctx.AbortWithStatus(http.StatusGone)
		return
	}

	_, err = c.DB.Exec("DELETE FROM CardioLog WHERE cardio_id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "cardio sucessfully deleted")
}

func (c CardioController) Create(ctx *gin.Context) {
	var cardio models.Cardio
	err := ctx.BindJSON(&cardio)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = c.DB.Exec("INSERT INTO Cardio (name) VALUES (?)", cardio.Name)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, "cardio sucessfully created")
}
