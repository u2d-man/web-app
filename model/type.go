package model

type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Username string
}