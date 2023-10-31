package dto

type UserRegister struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Deposit_Amount int    `json:"deposit_amount" binding:"required"`
}