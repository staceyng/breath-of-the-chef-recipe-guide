package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
