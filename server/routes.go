package server

import (
	"net/http"

	"github.com/labstack/echo"
)

// Routes: Defines Routes in this function
func Routes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}