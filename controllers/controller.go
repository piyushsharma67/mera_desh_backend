package controllers

import "social_web_server/services"


type ControllerStruct struct{
	service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.service=s

	return c
}