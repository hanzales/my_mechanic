package models

import (
	"time"
)

// Comment model
type Comment struct {
	Id        int       `json:"id" db:"id" validate:"omitempty"`
	Message   string    `json:"message" db:"message" validate:"required,gte=10"`
	Likes     int64     `json:"likes" db:"likes" validate:"omitempty"`
	UserId    int64     `json:"user_id" db:"user_id" validate:"omitempty"`
	Active    bool      `json:"active" db:"active"`
	DemandId  int64     `json:"demand_id" db:"demand_id" validate:"omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
