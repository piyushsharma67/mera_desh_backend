package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"social_web_server/database"
	"social_web_server/models"
	"social_web_server/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type GetPresignedUrlRequest struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
}

func (c *ControllerStruct) GetPresignedUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	var request models.GetPresignedUrl

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)

	defer cancel()
	fmt.Println(request)
	url, err := c.service.GetPresignedUrl(ctx, request.FileName, request.FileType, 2*time.Second)

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, url)
	return
}

func (c *ControllerStruct) SaveFileUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	var request models.FileUpload

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)

	defer cancel()

	err := c.service.SaveFileUrl(ctx, &database.CreateUploadedFileParams{
		UserID:  r.Context().Value("userid").(int32),
		FileUrl: request.FileUrl,
	})

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, nil)
	return
}
