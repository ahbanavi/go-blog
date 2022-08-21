package db

import (
	"fmt"

	"github.com/ahbanavi/go-blog/app/models"
	"github.com/ahbanavi/go-blog/pkg/common/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		env.GetD("DB_HOST", "127.0.0.1"),
		env.GetR("DB_USER"),
		env.GetR("DB_PASSWORD"),
		env.GetR("DB_NAME"),
		env.GetD("DB_PORT", "5432"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.User{})
}
