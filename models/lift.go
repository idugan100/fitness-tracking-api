package models

type Lift struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Compound bool   `json:"compound"`
	Upper    bool   `json:"upper"`
	Lower    bool   `json:"lower"`
}
