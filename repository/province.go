package repository

import (
	"github.com/sing3demons/thailand-provinces/entity"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	Find(zip_code string) ([]entity.ZipCode, error)
}

type provinceRepository struct{ DB *gorm.DB }

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepository{DB: db}
}

func (tx *provinceRepository) Find(zip_code string) ([]entity.ZipCode, error) {
	var provinces []entity.ZipCode

	if err := tx.DB.Where("zip_code = ?", zip_code).Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}
