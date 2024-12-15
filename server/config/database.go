package config

import (
	"fmt"
	"os"

	"tokutenban/models"
	"tokutenban/seeder"

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

func SeedDatabase(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE clubs")
	db.Exec("TRUNCATE TABLE formats")
	db.Exec("TRUNCATE TABLE individuals")
	db.Exec("TRUNCATE TABLE participants")
	db.Exec("TRUNCATE TABLE teams")
	db.Exec("TRUNCATE TABLE venues")
	db.Exec("TRUNCATE TABLE tournaments")
	db.Exec("TRUNCATE TABLE registrations")
	db.Exec("TRUNCATE TABLE face_offs")
	db.Exec("TRUNCATE TABLE rounds")
	db.Exec("TRUNCATE TABLE matches")
	db.Exec("TRUNCATE TABLE results")
	db.Exec("TRUNCATE TABLE shots")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	seeder.FormatSeeder(db)
	clubs := seeder.ClubSeeder(db)
	for _, club := range clubs {
		seeder.IndividualSeeder(db, seeder.IndividualOptions{Count: 10, Club: club})
		seeder.TeamSeeder(db, seeder.TeamOptions{Club: club, Size: 3})
		seeder.TeamSeeder(db, seeder.TeamOptions{Club: club, Size: 5})
	}

	venues := seeder.VenueSeeder(db, seeder.VenuOptions{Count: 10})
	for _, venue := range venues {
		for _, size := range []int{1, 3, 5} {
			var format models.Format
			db.Where("team_size = ?", size).First(&format)
			seeder.TournamentSeeder(db, seeder.TournamentOptions{Count: 1, Format: format, Venue: venue})
		}
	}
}
