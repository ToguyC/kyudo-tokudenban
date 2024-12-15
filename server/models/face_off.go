package models

type FaceOffFormat string

const (
	Closest  FaceOffFormat = "Closest"
	MostHits FaceOffFormat = "MostHits"
)

type FaceOff struct {
	ID     uint          `json:"id" gorm:"primaryKey"`
	Format FaceOffFormat `json:"format"`
}
