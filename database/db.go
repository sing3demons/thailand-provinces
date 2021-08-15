package database

import (
	"fmt"
	"os"

	"github.com/sing3demons/thailand-provinces/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func configDb() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, pass, name, port)
	return dsn
}

func InitDB() {
	dsn := configDb()

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(models.Amphue{})
	database.AutoMigrate(models.District{})
	database.AutoMigrate(models.Geographie{})
	database.AutoMigrate(models.Province{})
	database.AutoMigrate(models.ZipCode{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}
