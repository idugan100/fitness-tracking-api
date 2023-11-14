package models

import (
	"time"
)

type Workout struct {
	Id       int       `json:"id"`
	Location string    `json:"location" binding:"required"`
	Notes    string    `json:"notes" binding:"required"`
	Date     time.Time `json:"date"`
}
