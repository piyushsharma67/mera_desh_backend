package repository

import (
	"context"
	"social_web_server/database"
)


func (r *RepositoryStruct)GetUserByEmail(ctx context.Context,email string)(database.User,error){
	return r.db.GetUserByEmail(ctx,email)
}

func (r*RepositoryStruct)InsertUserInDB(ctx context.Context,details database.CreateUserParams)error{
	return r.db.CreateUser(ctx,details)
}