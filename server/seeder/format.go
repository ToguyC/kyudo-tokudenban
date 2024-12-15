package seeder

import (
	"log"
	"tokutenban/models"

	"gorm.io/gorm"
)

func FormatSeeder(db *gorm.DB) {
	log.Println("Seeding formats...")
	db.Save(&models.Format{Name: "Individual", TeamSize: 1})
	db.Save(&models.Format{Name: "Team", TeamSize: 2})
	db.Save(&models.Format{Name: "Team", TeamSize: 3})
	db.Save(&models.Format{Name: "Team", TeamSize: 4})
	db.Save(&models.Format{Name: "Team", TeamSize: 5})
}
