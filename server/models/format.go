package models

type Format struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	TeamSize int    `json:"team_size"`
}
