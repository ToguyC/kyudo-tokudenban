package models

type Participant struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	IndividualID *uint       `json:"individual_id"`
	Individual   *Individual `json:"individual" gorm:"foreignKey:IndividualID;references:ID"`
	TeamID       *uint       `json:"team_id"`
	Team         *Team       `json:"team" gorm:"foreignKey:TeamID;references:ID"`
}
