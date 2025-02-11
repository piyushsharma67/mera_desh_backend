package services

import (
	"context"
	"errors"
	"social_web_server/database"
	"social_web_server/models"
	"social_web_server/utils"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceStruct) InsertUserInDB(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := r.repository.GetUserByEmail(ctx, user.Email)

	if err != nil {
		// Check if error is "record not found"
		if errors.Is(err, pgx.ErrNoRows) {
			// User does not exist, so insert the user

			hashedPass, err := utils.HashPassword(user.Password)

			if err != nil {
				return nil, err
			}

			createUserParams := &database.CreateUserParams{
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
			}

			createUserParams.Password = hashedPass
			err = r.repository.InsertUserInDB(ctx, *createUserParams)
			if err != nil {
				return nil, err
			}

			// Fetch newly created user
			newUser, err := r.repository.GetUserByEmail(ctx, user.Email)
			if err != nil {
				return nil, err
			}

			user.ID = newUser.ID
			token, err := utils.EncodeToken(string(user.ID))

			if err != nil {
				return nil, utils.INTERNAL_SERVER_ERROR
			}

			user.Token = token

			return user, nil
		}

		// Return error if it's not a "not found" error
		return nil, err
	}

	return nil, utils.USER_ALREADY_EXISTS
}


func (r *ServiceStruct) InsertUserFCMInDB(ctx context.Context, user *models.UserFcm) (*models.User, error) {


	hashedPass, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	createUserParams := &database.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	createUserParams.Password = hashedPass
	err = r.repository.InsertUserInDB(ctx, *createUserParams)
	if err != nil {
		return nil, err
	}

	// Fetch newly created user
	newUser, err := r.repository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	user.ID = newUser.ID
	token, err := utils.EncodeToken(string(user.ID))

	if err != nil {
		return nil, utils.INTERNAL_SERVER_ERROR
	}

	user.Token = token

	return user, nil
}
	


