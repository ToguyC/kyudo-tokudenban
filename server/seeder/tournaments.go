package seeder

import (
	"log"
	"time"
	"tokutenban/models"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type TournamentOptions struct {
	Count  int
	Format models.Format
	Venue  models.Venue
}

func TournamentSeeder(db *gorm.DB, opt TournamentOptions) {
	log.Printf("Seeding %d tournaments of format %s(%d) in %s...", opt.Count, opt.Format.Name, opt.Format.TeamSize, opt.Venue.Name)

	for i := 0; i < opt.Count; i++ {
		d, _ := time.Parse("2006-01-02", faker.Date())
		tournament := models.Tournament{
			Name:     faker.Word(),
			Date:     d,
			Format:   opt.Format,
			FormatID: opt.Format.ID,
			Venue:    opt.Venue,
			VenueID:  opt.Venue.ID,
		}

		db.Create(&tournament)
	}
}
