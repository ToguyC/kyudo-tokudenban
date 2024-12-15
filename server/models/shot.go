package models

type HitState string

const (
	Hit     HitState = "Hit"
	Miss    HitState = "Miss"
	Unknown HitState = "Unknown"
)

type Shot struct {
	ID                 uint `json:"id" gorm:"primaryKey"`
	MatchID            uint
	Match              Match `json:"match" gorm:"foreignKey:MatchID;references:ID"`
	IndividualID       uint
	Individual         Individual `json:"individual" gorm:"foreignKey:IndividualID;references:ID"`
	ArrowNumber        uint       `json:"arrowNumber"`
	DistanceFromCenter float64    `json:"distanceFromCenter"`
	HitState           HitState   `json:"hitState"`
}
