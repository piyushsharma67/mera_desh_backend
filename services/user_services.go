package services

import (
	"context"
	"social_web_server/database"
)

type UserService struct{}

func (s *UserService) GetUserWithEmail(ctx context.Context,email string)(database.User, error){
	queries:=database.GetQueries()

	return queries.GetUserByEmail(ctx,email)
}