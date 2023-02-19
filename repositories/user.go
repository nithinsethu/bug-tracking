package repositories

import (
	"github.com/nithinsethu/bug-tracking/database/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(u *models.User) (*models.User, error) {
	result := ur.db.Create(u)
	return u, result.Error
}
