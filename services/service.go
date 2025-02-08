package services

import "social_web_server/db"

type Service struct{
	queries *db.Queries
}


func New