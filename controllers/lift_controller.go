package controllers

import (
	"fitness-tracker-api/testbackend/models"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLift(ctx *gin.Context) {
	//validate id input
	id, err := strconv.Atoi(ctx.Param("id"))
	//get result from database and read into struct

	row, err := db_connection.Query("SELECT * FROM LIFTS WHERE id=?", id)

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	defer row.Close()

	var lift models.Lift
	row_found := row.Next()

	if !row_found {
		ctx.AbortWithStatus(404)
		return
	}
	err = row.Scan(&lift.Id, &lift.Name, &lift.Compound, &lift.Upper, &lift.Lower)
	if err != nil {
		log.Print(row, err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSONP(200, lift)
	return
}

func GetAllLifts(ctx *gin.Context) {
	rows, err := db_connection.Query("SELECT * FROM LIFTS")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
	}
	defer rows.Close()

	var lifts []models.Lift
	for rows.Next() {
		var lift models.Lift
		err = rows.Scan(&lift.Id, &lift.Name, &lift.Compound, &lift.Upper, &lift.Lower)
		if err != nil {
			log.Print(lift, err)
			ctx.AbortWithStatus(500)
			return
		} else {
			lifts = append(lifts, lift)
		}
	}
	ctx.JSON(200, lifts)
}
