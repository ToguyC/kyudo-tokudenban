package models

type Team struct {
	ID      uint         `json:"id" gorm:"primaryKey"`
	Name    string       `json:"name"`
	Size    int          `json:"size"`
	Members []Individual `json:"members" gorm:"many2many:team_members;"`
}
