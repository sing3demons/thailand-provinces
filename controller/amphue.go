package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/sing3demons/thailand-provinces/service"
)

type Amphue interface {
	FindAmphue(c *fiber.Ctx) error
}

type amphue struct {
	Service service.AmphueService
}

func NewAmphue(service service.AmphueService) Amphue {
	return &amphue{Service: service}
}

type amphueResponse struct {
	Code       string `json:"code"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ProvinceID uint   `json:"province_id"`
	Province   struct {
		ID     uint   `json:"id"`
		Code   string `json:"code"`
		NameTH string `json:"name_th"`
		NameEN string `json:"name_en"`
	} `json:"provinces"`
}

func (tx *amphue) FindAmphue(c *fiber.Ctx) error {

	query := make(map[string]interface{})

	if code := c.Query("code"); code != "" {
		query["code"] = code
	}

	if nameTh := c.Query("name_th"); nameTh != "" {
		query["name_th"] = nameTh
	}

	if nameEn := c.Query("name_en"); nameEn != "" {
		query["name_en"] = nameEn
	}

	if provinceID := c.Query("province_id"); provinceID != "" {
		query["province_id"] = provinceID
	}

	amphues, err := tx.Service.Find(query)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	serializeAmphues := []amphueResponse{}
	copier.Copy(&serializeAmphues, &amphues)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"amphues": serializeAmphues})
}
