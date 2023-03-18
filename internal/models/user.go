package models

type User struct {
	ID uint `json:"id"`
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}