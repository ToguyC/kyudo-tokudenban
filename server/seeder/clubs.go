package seeder

import (
	"log"
	"tokutenban/models"

	"gorm.io/gorm"
)

func ClubSeeder(db *gorm.DB) []models.Club {
	log.Println("Seeding clubs...")
	clubs := []models.Club{
		{Name: "KKPLO"},
		{Name: "AKTBA"},
		{Name: "Kyudo Chablais"},
		{Name: "ALK"},
	}

	for i := 0; i < len(clubs); i++ {
		db.Create(&clubs[i])
	}

	return clubs
}
