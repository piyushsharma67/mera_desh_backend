package services

import (
	"context"
	"errors"
	"social_web_server/database"
	"social_web_server/utils"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceStruct) InsertUserInDB(ctx context.Context, user *database.CreateUserParams) (*database.User, error) {
	_, err := r.repository.GetUserByEmail(ctx, user.Email)

	if err != nil {
		// Check if error is "record not found"
		if errors.Is(err, pgx.ErrNoRows) {
			// User does not exist, so insert the user

			hashedPass,err := utils.HashPassword(user.Password)

			if err!=nil{
				return nil,err
			}

			user.Password=hashedPass
			err = r.repository.InsertUserInDB(ctx, *user)
			if err != nil {
				return nil, err
			}

			// Fetch newly created user
			newUser, err := r.repository.GetUserByEmail(ctx, user.Email)
			if err != nil {
				return nil, err
			}

			return &newUser, nil
		}

		// Return error if it's not a "not found" error
		return nil, err
	}

	return nil, utils.USER_ALREADY_EXISTS
}
