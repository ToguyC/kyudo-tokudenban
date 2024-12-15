package models

import "time"

type Tournament struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	FormatID uint
	VenueID  uint
	Format   Format `json:"format" gorm:"foreignKey:FormatID"`
	Venue    Venue  `json:"venue" gorm:"foreignKey:VenueID"`
}
