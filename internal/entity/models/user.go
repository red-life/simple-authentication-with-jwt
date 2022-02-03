package models

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primary_key:auto_increment"`
	Firstname string `gorm:"type:varchar(50)"`
	Lastname  string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(50)"`
	Username  string `gorm:"type:varchar(50)"`
	Password  string `gorm:"type:varchar(50)"`
	CreatedAt time.Time
}
