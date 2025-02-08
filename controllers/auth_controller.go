package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"social_web_server/models"
	"social_web_server/services"
	"social_web_server/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	userServiceInstance := &services.UserService{}

	db_user, err := userServiceInstance.GetUserWithEmail(ctx, user.Email)

	if db_user.Email != "" {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "user with the email already exists")
		return
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	dbUser, err := userServiceInstance.InserUser(ctx, user)

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = ""
	user.ID = dbUser.ID

	token, err := utils.EncodeToken(string(dbUser.ID))
	user.Token = token

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, user)

	return
}


func SigninUser(w http.ResponseWriter,r *http.Request){
	
}