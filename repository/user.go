package repository

import (
	"context"
	"social_web_server/db"
	"social_web_server/models"
	"time"
)

func InsertUser(ctx context.Context,user *models.User)error{
	query:=`INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.DB.QueryRow(ctx, query, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}