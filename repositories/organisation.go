package repositories

import (
	"github.com/nithinsethu/bug-tracking/database/models"
	"gorm.io/gorm"
)

type OrganisationRepository struct {
	db *gorm.DB
}

func NewOrganisationRepo(db *gorm.DB) *OrganisationRepository {
	return &OrganisationRepository{db: db}
}

func (or *OrganisationRepository) Save(o *models.Organisation) (*models.Organisation, error) {
	result := or.db.Create(o)

	return o, result.Error
}
