package seeder

import (
	"log"
	"tokutenban/models"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type TeamOptions struct {
	Club   models.Club
	Format models.Format
}

func TeamSeeder(db *gorm.DB, opt TeamOptions) {
	log.Println("Seeding teams...")

	var teamMembers []models.Individual
	db.Where("club_id = ?", opt.Club.ID).Limit(opt.Format.TeamSize).Find(&teamMembers)
	team := models.Team{
		Name:    faker.Word(),
		Members: teamMembers,
	}
	participant := models.Participant{
		Team: &team,
	}

	db.Create(&team)
	db.Create(&participant)
}
