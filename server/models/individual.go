package models

type Individual struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Dan           uint8  `json:"dan"`
	Title         string `json:"title"`
	ClubID        uint   `json:"club_id"`
	ParticipantID uint
	Participant   Participant `json:"participant" gorm:"foreignKey:ParticipantID;references:ID"`
}
