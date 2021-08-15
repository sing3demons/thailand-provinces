package repository

import (
	"github.com/sing3demons/thailand-provinces/models"
	"gorm.io/gorm"
)

type AmphueRepository interface {
	Find(search map[string]interface{}) ([]models.Amphue, error)
}

type amphueRepository struct {
	DB *gorm.DB
}

func NewAmphueRepository(db *gorm.DB) AmphueRepository {
	return &amphueRepository{DB: db}
}

func (tx *amphueRepository) Find(search map[string]interface{}) ([]models.Amphue, error) {

	var amphues []models.Amphue

	query := tx.DB.Preload("Province")

	if search["code"] != nil {
		query = query.Where("code = ?", search["code"])
	}

	if search["name_th"] != nil {
		query = query.Where("name_th = ?", search["name_th"])
	}

	if search["name_en"] != nil {
		query = query.Where("name_en = ?", search["name_en"])
	}

	if search["province_id"] != nil {
		query = query.Where("province_id = ?", search["province_id"])
	}

	if err := query.Find(&amphues).Error; err != nil {
		return nil, err
	}

	return amphues, nil
}
