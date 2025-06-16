package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/brunnossanttos/libri/internal/models"
)

var DB *gorm.DB

func Connect() error {
	dsn := viper.GetString("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	DB = db
	return nil
}
