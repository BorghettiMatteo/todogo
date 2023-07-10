package model

type User struct {
	Username       string `json:"username" gorm:"primaryKey" binding:"required"`
	HashedPassword string `json:"hashedpassword" binding:"required"`
}
