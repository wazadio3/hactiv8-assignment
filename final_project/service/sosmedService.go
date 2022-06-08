package service

import (
	"errors"
	"final_project/database"
	"final_project/models"
)

type SosMedService struct {
	sosMedQuery database.SosMedQuery
}

func NewSosMedService(query database.SosMedQuery) *SosMedService {
	return &SosMedService{
		sosMedQuery: query,
	}
}

func (s *SosMedService) CreateSosMed(sosMed models.SosMed) (models.SosMed, error) {
	res, err := s.sosMedQuery.CreateSosMed(sosMed)

	return res, err
}

func (s *SosMedService) GetSosMeds() ([]models.SosMed, error) {
	res, err := s.sosMedQuery.GetSosMeds()

	return res, err
}

func (s *SosMedService) UpdateSosMedById(userId uint, id interface{}, data models.SosMed) (models.SosMed, error) {
	sosMed, err := s.sosMedQuery.GetSosMedById(id)
	if err != nil {
		return sosMed, err
	}

	if sosMed.ID != userId {
		return sosMed, errors.New("Social media not found")
	}

	data.ID = sosMed.ID
	data.CreatedAt = sosMed.CreatedAt
	data.UserID = sosMed.UserID

	sosMed, err = s.sosMedQuery.UpdateSosMedById(data)
	if err != nil {
		return sosMed, err
	}

	return sosMed, err
}

func (s *SosMedService) DeleteSosMedById(userId uint, id interface{}) (string, error) {
	sosMed, err := s.sosMedQuery.GetSosMedById(id)
	if err != nil {
		return "", err
	}

	if sosMed.UserID != userId {
		return "", errors.New("Social media not found")
	}

	err = s.sosMedQuery.DeleteSosMedById(id)
	if err != nil {
		return "", err
	}

	return "Your social media deleted successfully", nil
}
