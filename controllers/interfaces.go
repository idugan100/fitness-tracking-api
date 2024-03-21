package controllers

import (
	"github.com/gin-gonic/gin"
)

type ExerciseController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Search(*gin.Context)
	Delete(*gin.Context)
	Create(*gin.Context)
}

type ExerciseLogController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	ByWorkout(*gin.Context)
	Delete(*gin.Context)
	Create(*gin.Context)
}

type ExerciseEventController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Delete(*gin.Context)
	Create(*gin.Context)
}
