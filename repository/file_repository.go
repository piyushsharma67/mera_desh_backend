package repository

import (
	"context"
	"social_web_server/database"
)

func (r *RepositoryStruct)InsertUploadedFile(ctx context.Context,params database.CreateUploadedFileParams)error{
	return r.db.CreateUploadedFile(ctx,params)
}