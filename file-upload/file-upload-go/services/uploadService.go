package services

import (
	"file-upload-go/domain"
	"mime/multipart"
)

type IUploadService interface {
	SaveFile(multipart.File, *multipart.FileHeader) (string, error)
}

type LocalUploadService struct {
	handler domain.IUpload
}

func (lus LocalUploadService) SaveFile(file multipart.File, handler *multipart.FileHeader) (string, error) {
	return lus.handler.SaveFile(file, handler)
}

func NewLocalUploadService(repo domain.IUpload) LocalUploadService {
	return LocalUploadService{
		handler: repo,
	}
}
