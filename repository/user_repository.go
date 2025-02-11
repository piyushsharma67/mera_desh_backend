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

func (r *RepositoryStruct)GetUserFcmById(ctx context.Context,userId int32)(database.UserFcmToken,error){
	return r.db.GetUserFcmTokenByUserID(ctx,userId)
}

func (r *RepositoryStruct)InsertUserFcmById(ctx context.Context,details database.CreateUserFcmTokenParams)(database.UserFcmToken,error){
	return r.db.CreateUserFcmToken(ctx,details)
}

func (r *RepositoryStruct)UpdateUserFcmById(ctx context.Context,details database.UpdateUserFcmTokenParams)(error){
	return r.db.UpdateUserFcmToken(ctx,details)
}