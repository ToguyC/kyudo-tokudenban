package seeder

import (
	"log"
	"tokutenban/models"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type VenuOptions struct {
	Count int
}

func VenueSeeder(db *gorm.DB, opt VenuOptions) []models.Venue {
	log.Println("Seeding venues...")

	var venues []models.Venue
	for i := 0; i < opt.Count; i++ {
		venue := models.Venue{
			Name:    faker.Word(),
			Address: faker.Sentence(),
		}
		db.Create(&venue)
		venues = append(venues, venue)
	}

	return venues
}
