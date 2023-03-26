package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nithinsethu/bug-tracking/enums"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email          string    `gorm:"unique"`
	Name           string
	OrganisationID uuid.UUID
	Organisation   Organisation
	Password       string
	Role           enums.Role `gorm:"type:role"`
	CreatedAt      time.Time
}

func (User) TableName() string {
	return "users"
}
