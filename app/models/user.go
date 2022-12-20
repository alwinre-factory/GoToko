package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Addresses     []Address //User Has Many Address
	FirstName     string    `gorm:"size:100;not null"`
	LastName      string    `gorm:"size:100;not null"`
	Email         string    `gorm:"size:100;not null:uniqueIndex"`
	Password      string    `gorm:"size:255;not null"`
	RememberToken string    `gorm:"size:255;not null"`
	CreateAt      time.Time
	UpdateAt      time.Time
	DeleteAt      gorm.DeletedAt
}
