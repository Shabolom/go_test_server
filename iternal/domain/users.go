package domain

type Users struct {
	Base
	Login    string `form:"login" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
