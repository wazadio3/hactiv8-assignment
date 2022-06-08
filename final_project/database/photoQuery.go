package database

import (
	"final_project/models"
)

func (p *DB) CreatePhoto(photo models.Photo) (models.Photo, error) {
	err := p.db.Create(&photo).Error
	return photo, err
}

func (p *DB) GetPhotos() ([]models.Photo, error) {
	photos := []models.Photo{}
	err := p.db.Find(&photos).Error

	return photos, err
}

func (p *DB) GetPhotoById(id interface{}) (models.Photo, error) {
	photo := models.Photo{}
	err := p.db.Where("id = ?", id).First(&photo).Error

	return photo, err
}

func (p *DB) UpdatePhotoById(photo models.Photo) (models.Photo, error) {
	err := p.db.Save(&photo).Error
	return photo, err
}

func (p *DB) DeletePhotoById(id interface{}) error {
	err := p.db.Delete(&models.Photo{}, id).Error
	return err
}
