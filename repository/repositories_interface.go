package repository

import (
	"context"
	"social_web_server/database"
	"social_web_server/models"
)

type UserRepository interface{
	GetUserByEmail(ctx context.Context,email string)(*models.User,error)
	InsertUserInDB(ctx context.Context,details database.CreateUserParams)error
	GetUserFcmById(ctx context.Context,userId int32)(*models.UserFcmToken,error)
	InsertUserFcmById(ctx context.Context,details database.CreateUserFcmTokenParams)(*models.UserFcmToken,error)
	UpdateUserFcmById(ctx context.Context,details database.UpdateUserFcmTokenParams)(error)
}

type FileUploadRepository interface {
	InsertUploadedFile(ctx context.Context,params database.CreateUploadedFileParams)error
	GetAllFilesOfUser(ctx context.Context, userId int32, limit int32, offset int32) ([]models.GetUserAllPhotosRow, error)
}