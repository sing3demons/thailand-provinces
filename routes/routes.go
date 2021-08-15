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
	amphueRepository := repository.NewAmphueRepository(db)
	amphueService := service.NewAmphueService(amphueRepository)
	amphueController := controller.NewAmphue(amphueService)
	amphueGroup := v1.Group("amphues")
	{
		amphueGroup.Get("", amphueController.FindAmphue)
	}

	districtRepository := repository.NewDistrictRepository(db)
	districtService := service.NewDistrictService(districtRepository)
	districtController := controller.NewDistrict(districtService)
	districtGroup := v1.Group("districts")

	{
		districtGroup.Get("", districtController.FindDistrict)
	}

	geographieGroup := v1.Group("geographies")
	geographieController := controller.NewGeographie(db)
	{
		geographieGroup.Get("", geographieController.FindGeographie)
	}

	provinceRepository:= repository.NewProvinceRepository(db)
	provinceService:=service.NewProvinceService(provinceRepository)
	provinceController := controller.NewProvince(provinceService)
	provinceGroup := v1.Group("provinces")
	{
		provinceGroup.Get("/:location_code", provinceController.FindProvince)
	}

}
