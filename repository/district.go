package repository

import (
	"github.com/sing3demons/thailand-provinces/models"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	Find(search map[string]interface{}) ([]models.District, error)
}

type districtRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtRepository{DB: db}
}

func (tx *districtRepository) Find(search map[string]interface{}) ([]models.District, error) {
	districts := []models.District{}

	query := tx.DB.Preload("Amphue")

	if search["zip_code"] != nil {
		query = query.Where("zip_code = ?", search["zip_code"])
	}

	if search["name_th"] != nil {
		query = query.Where("name_th = ?", search["name_th"])
	}

	if search["name_en"] != nil {
		query = query.Where("name_en = ?", search["name_en"])
	}

	if search["amphure_id"] != nil {
		query = query.Where("amphure_id = ?", search["amphure_id"])
	}

	if err := query.Find(&districts).Error; err != nil {
		return nil, err
	}

	return districts, nil
}
