package models

type GetPresignedUrl struct{
	FileName string `json:"file_name"`
	FileType string `json:"file_type"` 
}

type FileUpload struct{
	FileUrl string `json:"file_url" validate:"required"`
}

