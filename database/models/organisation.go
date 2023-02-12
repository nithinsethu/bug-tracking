package models

type Organisation struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
