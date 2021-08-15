package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/sing3demons/thailand-provinces/models"
	"gorm.io/gorm"
)

type Geographie interface {
	FindGeographie(c *fiber.Ctx) error
}

type geographie struct {
	DB *gorm.DB
}

type geographieResponse struct {
	Name string `json:"name"`
}

func NewGeographie(db *gorm.DB) Geographie {
	return &geographie{DB: db}
}

func (tx *geographie) FindGeographie(c *fiber.Ctx) error {
	var geographies []models.Geographie

	tx.DB.Find(&geographies)

	serializeGeographie := []geographieResponse{}
	copier.Copy(&serializeGeographie, &geographies)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"geographies": serializeGeographie})
}
