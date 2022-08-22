package database

import (
	"fmt"
	"log"

	"github.com/ahbanavi/go-blog/app/models"
	"github.com/ahbanavi/go-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.GetD("DB_HOST", "127.0.0.1"),
		config.GetR("DB_USER"),
		config.GetR("DB_PASSWORD"),
		config.GetR("DB_NAME"),
		config.GetD("DB_PORT", "5432"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(&models.User{})
}
