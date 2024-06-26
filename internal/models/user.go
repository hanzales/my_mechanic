package models

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User full model
type User struct {
	Id          int        `json:"id" db:"id"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Email       string     `json:"email,omitempty" db:"email"`
	Password    string     `json:"password,omitempty" db:"password"`
	Role        *string    `json:"role,omitempty" db:"role"`
	About       *string    `json:"about,omitempty" db:"about"`
	Avatar      *string    `json:"avatar,omitempty" db:"avatar"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number"`
	Address     *string    `json:"address,omitempty" db:"address"`
	City        *string    `json:"city,omitempty" db:"city"`
	Country     *string    `json:"country,omitempty" db:"country"`
	Gender      *string    `json:"gender,omitempty" db:"gender"`
	Postcode    *int       `json:"postcode,omitempty" db:"postcode"`
	Birthday    *time.Time `json:"birthday,omitempty" db:"birthday"`
	Active      bool       `json:"active" db:"active"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	LoginDate   time.Time  `json:"login_date" db:"login_date"`
}

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Role        *string
	About       *string
	Avatar      *string
	PhoneNumber *string
	Address     *string
	City        *string
	Country     *string
	Gender      *string
	Postcode    *int
	Birthday    *time.Time
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Prepare user for register
func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	if u.PhoneNumber != nil {
		*u.PhoneNumber = strings.TrimSpace(*u.PhoneNumber)
	}
	if u.Role != nil {
		*u.Role = strings.ToLower(strings.TrimSpace(*u.Role))
	}
	return nil
}

// Prepare user for register
func (u *User) PrepareUpdate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	if u.PhoneNumber != nil {
		*u.PhoneNumber = strings.TrimSpace(*u.PhoneNumber)
	}
	if u.Role != nil {
		*u.Role = strings.ToLower(strings.TrimSpace(*u.Role))
	}
	return nil
}

// All Users response
type UsersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}

// Find user query
type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
