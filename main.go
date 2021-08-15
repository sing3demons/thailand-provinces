package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sing3demons/thailand-provinces/database"
	"github.com/sing3demons/thailand-provinces/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	
	app := fiber.New()
	routes.Serve(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
