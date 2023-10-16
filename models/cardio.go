package models

type Cardio struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
