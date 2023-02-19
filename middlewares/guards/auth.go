package guards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nithinsethu/bug-tracking/constants"
	"github.com/nithinsethu/bug-tracking/services"
)

func AuthGuard(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil || !services.ValidateJWT(token) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constants.ErrorNotSignedIn})
		return
	}
	c.Next()
}
