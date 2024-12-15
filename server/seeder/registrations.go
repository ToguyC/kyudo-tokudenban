package seeder

import (
	"log"
	"tokutenban/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RegistrationOptions struct {
	Tournament  models.Tournament
	Participant models.Participant
}

func RegistrationSeeder(db *gorm.DB, opt RegistrationOptions) {
	log.Println("Seeding registrations...")

	registration := models.Registration{
		TournamentID:  opt.Tournament.ID,
		Tournament:    opt.Tournament,
		ParticipantID: opt.Participant.ID,
		Participant:   opt.Participant,
	}

	db.Omit(clause.Associations).Create(&registration)
}
