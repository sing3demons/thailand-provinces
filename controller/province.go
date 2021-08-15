package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/sing3demons/thailand-provinces/service"
)

type Province interface {
	FindProvince(c *fiber.Ctx) error
}

type province struct{ Service service.ProvinceService }

func NewProvince(service service.ProvinceService) Province {
	return &province{Service: service}
}

type provinceResponse struct {
	ZipCode     uint   `json:"zip_code"`
	Province    string `json:"province"`
	Amphoe      string `json:"amphoe"`
	Subdistrict string `json:"subdistrict"`
	Node        string `json:"node"`
}

func (tx *province) FindProvince(c *fiber.Ctx) error {
	zip_code := c.Params("location_code")
	provinces, err := tx.Service.Find(zip_code)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	serializeProvince := []provinceResponse{}
	copier.Copy(&serializeProvince, &provinces)

	return c.Status(fiber.StatusOK).JSON(serializeProvince)
}
