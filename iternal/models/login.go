package models

type Login struct {
	Login    string `form:"username" json:"login" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
