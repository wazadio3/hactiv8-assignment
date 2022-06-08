package service

import (
	"errors"
	"final_project/database"
	"final_project/models"
	"fmt"
)

type PhotoService struct {
	photoQuery database.PhotoQuery
}

func NewPhotoService(query database.PhotoQuery) *PhotoService {
	return &PhotoService{
		photoQuery: query,
	}
}

func (p *PhotoService) CreatePhoto(photo models.Photo) (models.Photo, error) {
	res, err := p.photoQuery.CreatePhoto(photo)
	return res, err
}

func (p *PhotoService) GetPhotos() ([]models.Photo, error) {
	res, err := p.photoQuery.GetPhotos()
	return res, err
}

func (p *PhotoService) UpdatePhotoById(userId uint, id interface{}, data models.Photo) (models.Photo, error) {
	photo, err := p.photoQuery.GetPhotoById(id)
	if err != nil {
		return photo, err
	}

	fmt.Println(photo.UserID, "=", userId)

	if photo.UserID != userId {
		return photo, errors.New("Photo not found")
	}

	data.ID = photo.ID
	data.CreatedAt = photo.CreatedAt
	data.UserID = photo.UserID
	photo, err = p.photoQuery.UpdatePhotoById(data)
	if err != nil {
		return photo, err
	}

	return photo, err
}

func (p *PhotoService) DeletePhotoById(userId uint, id interface{}) (string, error) {
	photo, err := p.photoQuery.GetPhotoById(id)
	if err != nil {
		return "", err
	}

	if photo.UserID != userId {
		return "", errors.New("Photo not found")
	}

	err = p.photoQuery.DeletePhotoById(id)
	if err != nil {
		return "", err
	}

	return "Your photo deleted successg=fully", nil
}
