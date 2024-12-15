package models

type Club struct {
	ID      uint         `json:"id" gorm:"primaryKey"`
	Name    string       `json:"name"`
	Members []Individual `json:"members" gorm:"foreignKey:ClubID"`
}
