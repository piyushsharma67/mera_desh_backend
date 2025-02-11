package services

import "social_web_server/database"

type ServiceStruct struct{
	db *database.Queries
}

func (s *ServiceStruct)InitialiseDB(queries *database.Queries){
	s.db = queries
}
