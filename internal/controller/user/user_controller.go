package controller

import (
	"net/http"

	domain "github.com/KOBATATU/todo/internal/domain/user"
	"github.com/KOBATATU/todo/internal/validation"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	requestBody := &domain.RegisterUser{}
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validation.ValidateStruct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, validation.GenerateValidationErrorMessages(err))
	}

	return c.String(200, "Register User")

}

func SetupUserRoutes(e *echo.Echo) {
	// userService := services.NewUserService()
	userController := NewUserController()

	e.POST("/user/signup", userController.RegisterUser)
}
