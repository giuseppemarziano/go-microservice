package entities

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Firstname string
	Surname   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
