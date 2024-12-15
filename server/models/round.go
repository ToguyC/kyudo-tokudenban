package models

type Round struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	TournamentID   uint
	Tournament     Tournament `json:"tournament" gorm:"foreignKey:TournamentID;references:ID"`
	Name           string     `json:"name"`
	SequenceNumber uint       `json:"sequenceNumber"`
}
