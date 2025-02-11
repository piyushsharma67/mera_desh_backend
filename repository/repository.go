package repository

import "social_web_server/database"

type RepositoryStruct struct{
	db *database.Queries
}

func (s *RepositoryStruct)InitialiseDB(queries *database.Queries)*RepositoryStruct{
	s.db = queries

	return s
}