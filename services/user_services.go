package services

import (
	"context"
	"social_web_server/database"
	"social_web_server/models"
)

type UserService struct{}

func (s *UserService) GetUserWithEmail(ctx context.Context,email string)(database.User, error){
	queries:=database.GetQueries()

	return queries.GetUserByEmail(ctx,email)
}

func (s *UserService)InserUser(ctx context.Context,user models.User)(database.User,error){
	queries:=database.GetQueries()
	params:=database.CreateUserParams{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
	
	err:=queries.CreateUser(ctx,params)

	if err !=nil{
		return database.User{},err
	}

	return s.GetUserWithEmail(ctx,user.Email)
}