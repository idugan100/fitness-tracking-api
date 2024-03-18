package models

type CardioLog struct {
	Id        int     `json:"id" binding:"required"`
	CardioId  int     `json:"cardio_id" binding:"required"`
	Time      int     `json:"time" binding:"required"`
	Distance  float64 `json:"distance" binding:"required"`
	WorkoutId int     `json:"workout_id" binding:"required"`
}
