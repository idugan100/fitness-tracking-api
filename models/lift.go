package models

type Lift struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Compound bool   `json:"compound" binding:"required"`
	Upper    bool   `json:"upper" binding:"required"`
	Lower    bool   `json:"lower" binding:"required"`
}
