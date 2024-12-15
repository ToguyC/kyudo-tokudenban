package models

type Team struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	Name          string       `json:"name"`
	ParticipantID uint         `json:"participant_id"`
	Members       []Individual `json:"members" gorm:"many2many:team_members;"`
}
