package main

import (
	controller "github.com/KOBATATU/todo/internal/controller/user"
	"github.com/KOBATATU/todo/internal/validation"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	validation.Init()

	controller.SetupUserRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
