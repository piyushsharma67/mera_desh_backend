package services

import (
	"context"
	"social_web_server/models"
	"social_web_server/repository"
)

func SaveUserInDb(ctx context.Context,user *models.User)error{

	if err:=repository.InsertUser(ctx,user);err!=nil{
		return err
	}

	return nil

}