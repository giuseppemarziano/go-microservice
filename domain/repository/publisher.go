package repository

import (
	"time"
)

type Publisher struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublisherRepository interface {
	Create(publisher *Publisher) error
	FindByID(id string) (*Publisher, error)
	FindAll(filter map[string]interface{}) ([]*Publisher, error)
	Update(id string, publisher *Publisher) error
	Delete(id string) error
}
