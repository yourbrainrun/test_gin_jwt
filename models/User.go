package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Region   string `gorm:"column:region"`
}
