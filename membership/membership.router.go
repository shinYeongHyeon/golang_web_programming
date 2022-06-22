package membership

import (
	"github.com/labstack/echo"
	membershipController "github.com/shinYeongHyeon/golang_web_programming/membership/controller"
	membershipPresentation "github.com/shinYeongHyeon/golang_web_programming/membership/presentation"
	membershipRepository "github.com/shinYeongHyeon/golang_web_programming/membership/repository"
	membershipServices "github.com/shinYeongHyeon/golang_web_programming/membership/services"
)

func getController() *membershipController.Controller {
	repository := membershipRepository.NewRepository(map[string]membershipPresentation.Membership{})
	service := membershipServices.NewService(*repository)
	controller := membershipController.NewController(*service)
	return controller
}

func CreateSubRouter(e *echo.Echo) {
	controller := getController()

	e.GET("/memberships", controller.FindAll)
	e.GET("/memberships/:id", controller.FindOne)
	e.POST("/memberships", controller.CreateMembership)
	e.PUT("/memberships/:id", controller.UpdateOne)
	e.DELETE("/memberships/:id", controller.DeleteOne)
}
