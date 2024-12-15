package seeder

import (
	"log"
	"tokutenban/models"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type TeamOptions struct {
	Club models.Club
	Size int
}

func TeamSeeder(db *gorm.DB, opt TeamOptions) {
	log.Println("Seeding teams...")

	var teamMembers []models.Individual
	db.Where("club_id = ?", opt.Club.ID).Limit(opt.Size).Find(&teamMembers)
	team := models.Team{
		Name:    faker.Word(),
		Size:    opt.Size,
		Members: teamMembers,
	}
	participant := models.Participant{
		Team: &team,
	}

	db.Create(&team)
	db.Create(&participant)
}
