package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetRecipes(c echo.Context) error {
	recipes := GetAllRecipes()
	return c.JSON(http.StatusOK, recipes)
}
