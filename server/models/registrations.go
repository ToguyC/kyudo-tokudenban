package models

type Registration struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	ParticipantID uint
	Participant   Participant `json:"participant" gorm:"foreignKey:ParticipantID;references:ID"`
	TournamentID  uint
	Tournament    Tournament `json:"tournament" gorm:"foreignKey:TournamentID;references:ID"`
}
