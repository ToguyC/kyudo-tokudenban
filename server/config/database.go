package config

import (
	"fmt"
	"math/rand"
	"os"

	"tokutenban/models"
	"tokutenban/seeder"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func randomOrder(n int) []int {
	// Create a slice of numbers from 0 to n-1
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	// Shuffle the slice
	rand.Shuffle(n, func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	return numbers
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

	venues := seeder.VenueSeeder(db, seeder.VenuOptions{Count: 3})
	for _, venue := range venues {
		for _, size := range []int{1, 3, 5} {
			var format models.Format
			db.Where("team_size = ?", size).First(&format)
			seeder.TournamentSeeder(db, seeder.TournamentOptions{Count: 1, Format: format, Venue: venue})
		}
	}

	var tournaments []models.Tournament
	db.Preload(clause.Associations).Find(&tournaments)
	db.Preload(clause.Associations).Find(&clubs)

	for _, tournament := range tournaments {
		if tournament.Format.TeamSize == 1 {
			for _, i := range randomOrder(10) {
				var participant models.Participant
				db.Where("individual_id = ?", i).First(&participant)
				seeder.RegistrationSeeder(db, seeder.RegistrationOptions{Tournament: tournament, Participant: participant})
			}
		} else {
			var teams []models.Team
			db.Preload(clause.Associations).Where("size = ?", tournament.Format.TeamSize).Find(&teams)
			for _, team := range teams {
				var participant models.Participant
				db.Where("team_id = ?", team.ID).First(&participant)
				seeder.RegistrationSeeder(db, seeder.RegistrationOptions{Tournament: tournament, Participant: participant})
			}
		}
	}
}
