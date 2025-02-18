package repository

import (
	"context"
	"social_web_server/database"
	"social_web_server/models"
)

type PostgresFileRepository struct {
	db *database.Queries
}

func NewPostgresFileRepository(db *database.Queries) *PostgresFileRepository {
	return &PostgresFileRepository{db: db}
}

func (r *PostgresFileRepository) InsertUploadedFile(ctx context.Context, params database.CreateUploadedFileParams) error {
	return r.db.CreateUploadedFile(ctx, params)
}

func (r *PostgresFileRepository) GetAllFilesOfUser(ctx context.Context, userId int32, limit int32, offset int32) ([]models.GetUserAllPhotosRow, error) {
	userFile,err:=r.db.GetUserAllPhotos(ctx, database.GetUserAllPhotosParams{
		UserID: userId,
		Limit:  limit,
		Offset: offset,
	})

	if err !=nil{
		return nil,err
	}

	return ConvertToModelPhotos(userFile),nil
}

func ConvertToModelPhotos(dbPhotos []database.GetUserAllPhotosRow) []models.GetUserAllPhotosRow {
    modelPhotos := make([]models.GetUserAllPhotosRow, len(dbPhotos))
    for i, photo := range dbPhotos {
        modelPhotos[i] = models.GetUserAllPhotosRow{
            ID:        photo.ID,
            UserID:    photo.UserID,
            FileUrl:  photo.FileUrl,
            CreatedAt: photo.CreatedAt,
        }
    }
    return modelPhotos
}

