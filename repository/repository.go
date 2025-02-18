package repository

import (
	"social_web_server/database"
	"social_web_server/enums"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct{
	UserRepo UserRepository
	FileUploadRepo FileUploadRepository
}

func InitialiseRepositories(dbType enums.DBType,postgresDB *database.Queries,mongoClient *mongo.Client)*Repositories{
	var userRepo UserRepository
	var fileUpload FileUploadRepository

	switch dbType{

	case enums.Postgres:
		userRepo=NewPostgresUserRepository(postgresDB)
		fileUpload=NewPostgresFileRepository(postgresDB)
	}

	return &Repositories{
		UserRepo: userRepo,
		FileUploadRepo: fileUpload,
	}
}