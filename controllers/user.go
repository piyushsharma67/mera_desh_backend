package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"social_web_server/database"
	"social_web_server/models"
	"social_web_server/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func (c *ControllerStruct)SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*3)

	defer cancel()

	createUserParams:=&database.CreateUserParams{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	_,err:=c.service.InsertUserInDB(ctx,createUserParams)

	if err!=nil{
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	
}
