package config

import (
	"fmt"
	"os"

	"tokutenban/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Club{})
	db.AutoMigrate(&models.Format{})
	db.AutoMigrate(&models.Individual{})
	db.AutoMigrate(&models.Participant{})
	db.AutoMigrate(&models.Team{})
	db.AutoMigrate(&models.Venue{})
	db.AutoMigrate(&models.Tournament{})
	db.AutoMigrate(&models.Registration{})
	db.AutoMigrate(&models.FaceOff{})
	db.AutoMigrate(&models.Round{})
	db.AutoMigrate(&models.Match{})
	db.AutoMigrate(&models.Result{})
	db.AutoMigrate(&models.Shot{})
}
