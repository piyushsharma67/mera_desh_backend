package services

import "social_web_server/repository"

type ServiceStruct struct{
	repository *repository.RepositoryStruct
}

func (s *ServiceStruct)InitialiseService(r *repository.RepositoryStruct)(*ServiceStruct){
	s.repository=r

	return s
}
