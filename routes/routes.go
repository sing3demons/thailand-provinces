package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/thailand-provinces/controller"
	"github.com/sing3demons/thailand-provinces/database"
	"github.com/sing3demons/thailand-provinces/repository"
	"github.com/sing3demons/thailand-provinces/service"
)

func Serve(app *fiber.App) {
	db := database.GetDB()
	v1 := app.Group("/api/v1/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello, world")
	})

	provinceRepository := repository.NewProvinceRepository(db)
	provinceService := service.NewProvinceService(provinceRepository)
	provinceController := controller.NewProvince(provinceService)
	provinceGroup := v1.Group("provinces")
	{
		provinceGroup.Get("/:location_code", provinceController.FindProvince)
	}

}
