package models

import "time"

type User struct {
	ID        int32   `json:"id"`
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Token string `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFcm struct {
	ID        int32   `json:"id"`
	FcmToken string `json:"fcm_token"`
	CreatedAt time.Time `json:"created_at"`
}