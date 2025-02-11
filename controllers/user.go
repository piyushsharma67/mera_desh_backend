package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"social_web_server/models"
	"social_web_server/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func (c *ControllerStruct) SignupUser(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)

	defer cancel()

	db_user, err := c.service.InsertUserInDB(ctx, &user)

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, db_user)
	return
}

func (c *ControllerStruct)SaveUserFcmToken(w http.ResponseWriter,r *http.Request){
	userId,ok:=r.Context().Value("userid").(int32)

	if !ok{
		utils.ErrorResponse(w, r, http.StatusForbidden, utils.INTERNAL_SERVER_ERROR.Error())
		return
	}

	
}
