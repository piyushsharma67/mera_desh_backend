package repository

import "go.mongodb.org/mongo-driver/mongo"

type MongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db: db}
}

type MongoFileUploadRepository struct {
	db *mongo.Database
}

func NewMongoFileUploadRepository(db *mongo.Database) *MongoFileUploadRepository {
	return &MongoFileUploadRepository{db: db}
}