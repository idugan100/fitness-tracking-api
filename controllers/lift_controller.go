package controllers

import (
	"database/sql"
	"fitness-tracker-api/testbackend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LiftController struct {
	DB *sql.DB
}

func NewLiftController(DB *sql.DB) *LiftController {
	return &LiftController{DB}
}

func (lc *LiftController) GetLift(ctx *gin.Context) {
	//validate id input
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "invalid id parameter")
		return
	}

	//start repo
	row, err := lc.DB.Query("SELECT * FROM LIFTS WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var lift models.Lift
	row_found := row.Next()

	if !row_found {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = row.Scan(&lift.Id, &lift.Name, &lift.Compound, &lift.Upper, &lift.Lower)
	if err != nil {
		log.Print(row, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//end repo
	ctx.JSON(http.StatusOK, lift)
}

func (lc *LiftController) GetAllLifts(ctx *gin.Context) {
	//start repo
	rows, err := lc.DB.Query("SELECT * FROM LIFTS")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	defer rows.Close()

	var lifts []models.Lift
	var lift models.Lift
	for rows.Next() {
		err = rows.Scan(&lift.Id, &lift.Name, &lift.Compound, &lift.Upper, &lift.Lower)
		if err != nil {
			log.Print(lift, err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			lifts = append(lifts, lift)
		}
	}
	//end repo
	ctx.JSON(http.StatusOK, lifts)
}

func (lc *LiftController) SearchLiftsByName(ctx *gin.Context) {
	search_name := ctx.Query("name")
	if search_name == "" {
		ctx.JSON(http.StatusBadRequest, "missing name query parameter")
		return
	}
	//start repo
	rows, err := lc.DB.Query("SELECT * FROM Lifts WHERE name like ?", "%"+search_name+"%")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lifts []models.Lift
	for rows.Next() {
		var lift models.Lift
		err = rows.Scan(&lift.Id, &lift.Name, &lift.Compound, &lift.Upper, &lift.Lower)
		log.Print(lift)
		if err != nil {
			log.Print(lift, err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			lifts = append(lifts, lift)
		}
	}
	//end repo
	ctx.JSON(http.StatusOK, lifts)
}

func (lc *LiftController) DeleteLift(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//start repo
	res, err := lc.DB.Exec("DELETE FROM Lifts WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows_deleted, err := res.RowsAffected()
	//end repo
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if rows_deleted == 0 {
		log.Print("Unsuccessful Deletion - Resource Not Found")
		ctx.AbortWithStatus(http.StatusGone)
		return
	}
	_, err = lc.DB.Exec("DELETE FROM LiftingLog WHERE lift_id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "lift successfully deleted")

}

func (lc *LiftController) AddLift(ctx *gin.Context) {
	var new_lift models.Lift
	err := ctx.BindJSON(&new_lift)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//start repo
	_, err = lc.DB.Exec("INSERT INTO Lifts (name, compound, upper, lower) Values (?,?,?,?) ", new_lift.Name, new_lift.Compound, new_lift.Upper, new_lift.Lower)
	//end repo
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, "lift sucessfully created")

}
