package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null;" json:"firstName"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	LastName  string `gorm:"size:255;not null;" json:"lastName"`
	Email     string `gorm:"email:varchar(120);unique"`
}
