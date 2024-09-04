package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey;autoIncrement:false"`
	Email    string
	Password string
}
