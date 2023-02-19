package models

import "github.com/google/uuid"

type Organisation struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string
}

func (Organisation) TableName() string {
	return "organisations"
}
