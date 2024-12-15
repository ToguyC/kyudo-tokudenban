package models

type Result struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	TournamentID  uint
	Tournament    Tournament `json:"tournament" gorm:"foreignKey:TournamentID;references:ID"`
	FirstPlaceID  uint
	FirstPlace    Participant `json:"firstPlace" gorm:"foreignKey:FirstPlaceID;references:ID"`
	SecondPlaceID uint
	SecondPlace   Participant `json:"secondPlace" gorm:"foreignKey:SecondPlaceID;references:ID"`
	ThirdPlaceID  uint
	ThirdPlace    Participant `json:"thirdPlace" gorm:"foreignKey:ThirdPlaceID;references:ID"`
	FourthPlaceID uint
	FourthPlace   Participant `json:"fourthPlace" gorm:"foreignKey:FourthPlaceID;references:ID"`
	FirstMatchID  uint
	FirstMatch    Match `json:"firstMatch" gorm:"foreignKey:FirstMatchID;references:ID"`
	SecondMatchID uint
	SecondMatch   Match `json:"secondMatch" gorm:"foreignKey:SecondMatchID;references:ID"`
}
