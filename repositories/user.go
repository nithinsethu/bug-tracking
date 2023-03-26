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

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := ur.db.Where("email = ?", email).First(&user)
	return user, result.Error
}
