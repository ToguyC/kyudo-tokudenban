package models

type Participant struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	FormatID uint
	Format   Format `json:"format" gorm:"foreignKey:FormatID;references:ID"`
}
