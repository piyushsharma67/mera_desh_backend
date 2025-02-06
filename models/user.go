package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
}