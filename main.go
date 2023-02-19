package main

import (
	"net/http"

	"github.com/nithinsethu/bug-tracking/constants"
	"github.com/nithinsethu/bug-tracking/controllers"
	"github.com/nithinsethu/bug-tracking/database"
	"github.com/nithinsethu/bug-tracking/repositories"
	"github.com/nithinsethu/bug-tracking/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//***Initializations***

	//Database
	pg := database.NewPostgresDB()
	pg.InitDB()

	//Repos
	or := repositories.NewOrganisationRepo(pg.GetdbInstance())
	ur := repositories.NewUserRepo(pg.GetdbInstance())

	//Services
	as := services.NewAuthService(ur, or)

	//Auth Controller
	ac := controllers.NewAuthController(as)

	//Router
	r := gin.Default()

	//***Registering Routes***

	rootGroup := r.Group(constants.RouteRoot)
	rootGroup.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Auth routes
	ac.AddRoutes(rootGroup)

	r.Run()
}
