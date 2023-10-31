package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getAllWorkouts(ctx *gin.Context) {

	rowzzz, err := db_connection.Query("SELECT * FROM Workouts")

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(500)
		return
	}

	return ctx.JSON(200)
}
