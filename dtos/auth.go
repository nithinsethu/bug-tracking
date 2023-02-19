package dtos

type SignupRequest struct {
	Email            string `json:"email" binding:"required,email"`
	Name             string `json:"name" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganisationName string `json:"organisationName" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
