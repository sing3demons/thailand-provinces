package service

import (
	"log"

	"github.com/sing3demons/thailand-provinces/models"
	"github.com/sing3demons/thailand-provinces/repository"
)

type AmphueService interface {
	Find(search map[string]interface{}) ([]models.Amphue, error)
}

type amphueService struct {
	Repository repository.AmphueRepository
}

func NewAmphueService(repository repository.AmphueRepository) AmphueService {
	return &amphueService{Repository: repository}
}

func (service *amphueService) Find(search map[string]interface{}) ([]models.Amphue, error) {
	amphues, err := service.Repository.Find(search)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return amphues, nil
}
