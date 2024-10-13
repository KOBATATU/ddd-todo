package controller

import (
	"context"
	"net/http"

	domain "github.com/KOBATATU/todo/internal/domain/user"
	service "github.com/KOBATATU/todo/internal/service/user"
	"github.com/KOBATATU/todo/internal/validation"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	requestBody := &domain.RegisterUser{}
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validation.ValidateStruct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, validation.GenerateValidationErrorMessages(err))
	}

	ctx := context.Background()
	user, err := uc.service.CreateUser(ctx, requestBody.Email, requestBody.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, user)

}

func SetupUserRoutes(e *echo.Echo, service *service.UserService) {
	userController := NewUserController(service)
	e.POST("/user/signup", userController.RegisterUser)
}
