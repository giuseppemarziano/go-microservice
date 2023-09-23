package entities

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Post) TableName() string {
	return "posts"
}
