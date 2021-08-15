package repository

import (
	"github.com/sing3demons/thailand-provinces/models"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	Find(zip_code string) ([]models.ZipCode, error)
}

type provinceRepository struct{ DB *gorm.DB }

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepository{DB: db}
}

func (tx *provinceRepository) Find(zip_code string) ([]models.ZipCode, error) {
	var provinces []models.ZipCode

	if err := tx.DB.Where("zip_code = ?", zip_code).Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}
