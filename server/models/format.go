package models

type Format struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
