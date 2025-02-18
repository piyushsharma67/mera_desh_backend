package repository

import (
	"context"
	"social_web_server/database"
	"social_web_server/models"
)

type PostgresUserRepository struct {
	db *database.Queries
}

func NewPostgresUserRepository(db *database.Queries) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Token:     "",
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *PostgresUserRepository) InsertUserInDB(ctx context.Context, details database.CreateUserParams) error {
	return r.db.CreateUser(ctx, details)
}

func (r *PostgresUserRepository) GetUserFcmById(ctx context.Context, userId int32) (*models.UserFcmToken, error) {
	userFcm, err := r.db.GetUserFcmTokenByUserID(ctx, userId)

	if err != nil {
		return nil, err
	}
	return &models.UserFcmToken{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresUserRepository) InsertUserFcmById(ctx context.Context, details database.CreateUserFcmTokenParams) (*models.UserFcmToken, error) {
	userFcm, err := r.db.CreateUserFcmToken(ctx, details)

	if err != nil {
		return nil, err
	}
	return &models.UserFcmToken{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresUserRepository) UpdateUserFcmById(ctx context.Context, details database.UpdateUserFcmTokenParams) error {
	return r.db.UpdateUserFcmToken(ctx, details)
}

func (r *PostgresUserRepository) GetAllUserUploadedFiles(ctx context.Context, userId int32, limit int32, offset int32) ([]models.GetUserAllPhotosRow, error) {

	userUploadedFile, err := r.db.GetUserAllPhotos(ctx, database.GetUserAllPhotosParams{
		UserID: userId,
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	return mapDatabasePhotosToModel(userUploadedFile), nil
}

func mapDatabasePhotosToModel(dbPhotos []database.GetUserAllPhotosRow) []models.GetUserAllPhotosRow {
	var modelPhotos []models.GetUserAllPhotosRow

	for _, photo := range dbPhotos {
		modelPhotos = append(modelPhotos, models.GetUserAllPhotosRow{
			ID:        photo.ID,
			UserID:    photo.UserID,
			FileUrl:   photo.FileUrl,
			CreatedAt: photo.CreatedAt,
		})
	}

	return modelPhotos
}
