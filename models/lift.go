package models

// using *bool due to issue with gin validators here: https://github.com/gin-gonic/gin/issues/814#issuecomment-294636138
type Lift struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Compound *bool  `json:"compound" binding:"required"`
	Upper    *bool  `json:"upper" binding:"required"`
	Lower    *bool  `json:"lower" binding:"required"`
}
