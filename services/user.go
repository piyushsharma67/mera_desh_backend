package services

import (
	"context"
	"social_web_server/db"
	"social_web_server/models"
)



func SaveUserInDb(ctx context.Context,user *models.User)error{
	db.New().CreateUser()
}