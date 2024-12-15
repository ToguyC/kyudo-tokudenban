package models

type Match struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	RoundID      uint
	Round        Round `json:"round" gorm:"foreignKey:RoundID;references:ID"`
	WinnerID     uint
	Winner       *Participant `json:"winner" gorm:"foreignKey:WinnerID;references:ID"`
	FaceOffID    uint
	FaceOff      *Participant  `json:"faceOff" gorm:"foreignKey:FaceOffID;references:ID"`
	Participants []Participant `json:"participants" gorm:"many2many:match_participants;"`
}
