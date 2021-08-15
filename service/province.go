package service

import (
	"log"

	"github.com/sing3demons/thailand-provinces/models"
	"github.com/sing3demons/thailand-provinces/repository"
)

type ProvinceService interface {
	Find(zip_code string) ([]models.ZipCode, error)
}

type provinceService struct{ Repository repository.ProvinceRepository }

func NewProvinceService(repository repository.ProvinceRepository) ProvinceService {
	return &provinceService{Repository: repository}
}

func (service provinceService) Find(zip_code string) ([]models.ZipCode, error) {
	provinces, err := service.Repository.Find(zip_code)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return provinces, nil
}
