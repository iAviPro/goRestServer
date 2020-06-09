package server

import (
	"github.com/iAviPro/goRestSetup/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start Server on the Port: 1220
func Start() {
	e := echo.New()
	Routes(e)
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(config.PORT))
}
