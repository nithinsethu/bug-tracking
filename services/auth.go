package services

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nithinsethu/bug-tracking/config"
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

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
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

func (as *AuthService) LoginUser(r dtos.LoginRequest) (*http.Cookie, error) {
	u, err := as.ur.FindByEmail(r.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New(constants.ErrorInvalidEmail)
	}
	if !isPasswordValid(u.Password, r.Password) {
		return nil, errors.New(constants.ErrorInvalidPassword)
	}
	expiry := time.Now().Add(24 * time.Hour)
	token, err := generateJWT(r.Email, expiry)
	if err != nil {
		log.Println(err)
		return nil, errors.New(constants.ErrorUnknown)
	}

	return &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiry,
		Path:    constants.RouteRoot,
	}, nil
}

func hashAndSaltPassword(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(bytes)
}

func isPasswordValid(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func generateJWT(email string, expiry time.Time) (string, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ValidateJWT(tokenString string) bool {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		log.Println(err)
		return false
	}
	return token.Valid
}
