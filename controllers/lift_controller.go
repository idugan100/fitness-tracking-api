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
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(500, "invalid id parameter")
		return
	}
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

	ctx.JSON(200, lift)
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

func SearchLiftsByName(ctx *gin.Context) {
	search_name := ctx.Query("name")
	if search_name == "" {
		ctx.JSON(400, "missing name query parameter")
		return
	}
	rows, err := db_connection.Query("SELECT * FROM Lifts WHERE name like ?", "%"+search_name+"%")
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
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
			ctx.AbortWithStatus(500)
			return
		} else {
			lifts = append(lifts, lift)
		}
	}
	ctx.JSON(200, lifts)
}

func DeleteLift(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	_, err = db_connection.Exec("DELETE FROM Lifts WHERE id=?", id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(200, "lift sucessfully deleted")

}

func AddLift(ctx *gin.Context) {
	var new_lift models.Lift
	err := ctx.BindJSON(&new_lift)
	if err != nil {
		log.Print(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = db_connection.Exec("INSERT INTO Lifts (name, compound, upper, lower) Values (?,?,?,?) ", new_lift.Name, new_lift.Compound, new_lift.Upper, new_lift.Lower)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSON(201, "lift sucessfully created")

}
