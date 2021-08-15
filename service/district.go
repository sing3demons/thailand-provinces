package service

import (
	"log"

	"github.com/sing3demons/thailand-provinces/models"
	"github.com/sing3demons/thailand-provinces/repository"
)

type DistrictService interface {
	Find(map[string]interface{}) ([]models.District, error)
}

type districtService struct {
	Repository repository.DistrictRepository
}

func NewDistrictService(repository repository.DistrictRepository) DistrictService {
	return &districtService{Repository: repository}
}

func (service *districtService) Find(search map[string]interface{}) ([]models.District, error) {
	districts, err := service.Repository.Find(search)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return districts, nil
}
