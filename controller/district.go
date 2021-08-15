package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/sing3demons/thailand-provinces/service"
)

type District interface {
	FindDistrict(c *fiber.Ctx) error
}

type district struct {
	Service service.DistrictService
}

func NewDistrict(service service.DistrictService) District {
	return &district{Service: service}
}

type districtResponse struct {
	ZipCode   uint   `json:"zip_code"`
	NameTH    string `json:"name_th"`
	NameEN    string `json:"name_en"`
	AmphureID uint   `json:"amphure_id"`
	Amphue    struct {
		Code   string `json:"code"`
		NameTH string `json:"name_th"`
		NameEN string `json:"name_en"`
	} `json:"amphure"`
}

type response struct {
	ZipCode   uint `json:"zip_code"`
	AmphureID uint `json:"amphure_id"`
	Amphue    struct {
		ProvinceID uint `json:"province_id"`
	} `json:"amphure"`
}

func (tx *district) FindDistrict(c *fiber.Ctx) error {
	query := make(map[string]interface{})

	if zipCode := c.Query("zip_code"); zipCode != "" {
		query["zip_code"] = zipCode
	}

	if nameTh := c.Query("name_th"); nameTh != "" {
		query["name_th"] = nameTh
	}

	if nameEn := c.Query("name_en"); nameEn != "" {
		query["name_en"] = nameEn
	}

	if AmphureID := c.Query("amphure_id"); AmphureID != "" {
		query["amphure_id"] = AmphureID
	}

	districts, err := tx.Service.Find(query)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if len(query) == 0 {
		serializeDistrict := []response{}
		copier.Copy(&serializeDistrict, &districts)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"districts": serializeDistrict})
	}

	serializeDistrict := []districtResponse{}
	copier.Copy(&serializeDistrict, &districts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"districts": serializeDistrict})
}
