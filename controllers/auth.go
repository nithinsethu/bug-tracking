package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nithinsethu/bug-tracking/constants"
	"github.com/nithinsethu/bug-tracking/dtos"
	"github.com/nithinsethu/bug-tracking/services"
)

type AuthController struct {
	as *services.AuthService
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{as: as}
}

func (ac *AuthController) AddRoutes(rootGroup *gin.RouterGroup) {
	authGroup := rootGroup.Group(constants.RouteAuth)
	authGroup.POST(constants.RouteLogin, ac.loginHandler)
	authGroup.POST(constants.RouteSignup, ac.signupHandler)
}

func (ac *AuthController) signupHandler(c *gin.Context) {
	var sr dtos.SignupRequest
	err := c.BindJSON(&sr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = ac.as.SignupUser(sr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Status(http.StatusCreated)
}

func (ac *AuthController) loginHandler(c *gin.Context) {
	var lr dtos.LoginRequest
	err := c.BindJSON(&lr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	cookie, err := ac.as.LoginUser(lr)
	if err != nil {
		if err.Error() == constants.ErrorUnknown {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	http.SetCookie(c.Writer, cookie)
	c.Status(http.StatusOK)
}
