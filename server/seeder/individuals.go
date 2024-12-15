package seeder

import (
	"log"
	"math/rand"
	"tokutenban/models"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IndividualOptions struct {
	Count int
	Club  models.Club
}

var titles = []string{
	"Renshi",
	"Kyoshi",
	"Hanshi",
}

func generateDan() int {
	weights := []int{10, 10, 20, 35, 25, 10, 5, 1, 1}
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}

	threshold := rand.Intn(totalWeight)

	currentSum := 0
	for i, weight := range weights {
		currentSum += weight
		if currentSum >= threshold {
			return i
		}
	}

	return 0
}

func generateTitle(dan int) *string {
	threshold := rand.Intn(100)
	if dan < 5 {
		return nil
	}

	switch dan {
	case 5:
		if threshold < 10 {
			return &titles[0]
		}
	case 6:
		if threshold < 5 {
			return nil
		} else if threshold < 20 {
			return &titles[0]
		} else if threshold < 50 {
			return &titles[1]
		}
	case 7:
		return &titles[1]
	case 8:
		return &titles[2]
	}

	return nil
}

func IndividualSeeder(db *gorm.DB, opt IndividualOptions) {
	log.Printf("Seeding %d individuals for club %s...\n", opt.Count, opt.Club.Name)
	for i := 0; i < opt.Count; i++ {
		dan := generateDan()
		user := models.Individual{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Dan:       uint8(dan),
			Title:     generateTitle(dan),
			ClubID:    opt.Club.ID,
			Club:      opt.Club,
		}
		participant := models.Participant{
			Individual: &user,
		}

		db.Omit(clause.Associations).Create(&user)
		db.Create(&participant)
	}
}
