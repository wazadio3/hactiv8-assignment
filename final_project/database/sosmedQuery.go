package database

import "final_project/models"

func (s *DB) CreateSosMed(sosMed models.SosMed) (models.SosMed, error) {
	err := s.db.Create(&sosMed).Error

	return sosMed, err
}

func (s *DB) GetSosMeds() ([]models.SosMed, error) {
	sosMeds := []models.SosMed{}
	err := s.db.Find(&sosMeds).Error

	return sosMeds, err
}

func (s *DB) GetSosMedById(id interface{}) (models.SosMed, error) {
	sosMed := models.SosMed{}
	err := s.db.Where("id = ?", id).First(&sosMed).Error

	return sosMed, err
}

func (s *DB) UpdateSosMedById(sosMed models.SosMed) (models.SosMed, error) {
	err := s.db.Save(&sosMed).Error

	return sosMed, err
}

func (s *DB) DeleteSosMedById(id interface{}) error {
	err := s.db.Delete(&models.SosMed{}, id).Error

	return err
}
