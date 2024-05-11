package models

import (
	"time"
)

// Comment model
type Comment struct {
	Id        int       `json:"id" db:"id" validate:"omitempty"`
	Message   string    `json:"message" db:"message" validate:"required,gte=10"`
	Likes     int64     `json:"likes" db:"likes" validate:"omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
