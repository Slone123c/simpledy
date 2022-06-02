package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:32;not null"`
	Password string `gorm:"size:32;not null"`
}
