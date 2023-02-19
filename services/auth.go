package services

import (
	"log"

	"github.com/nithinsethu/bug-tracking/constants"
	"github.com/nithinsethu/bug-tracking/database/models"
	"github.com/nithinsethu/bug-tracking/dtos"
	"github.com/nithinsethu/bug-tracking/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	ur *repositories.UserRepository
	or *repositories.OrganisationRepository
}

func NewAuthService(ur *repositories.UserRepository, or *repositories.OrganisationRepository) *AuthService {
	return &AuthService{ur: ur, or: or}
}

func (as *AuthService) SignupUser(r dtos.SignupRequest) error {

	o, err := as.or.Save(&models.Organisation{Name: r.OrganisationName})
	if err != nil {
		log.Println(err)
		return err
	}
	hp := hashAndSaltPassword(r.Password)
	_, err = as.ur.Save(&models.User{Email: r.Email, Name: r.Name, OrganisationID: o.ID, Role: constants.RoleAdmin, Password: hp})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func hashAndSaltPassword(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(bytes)
}
