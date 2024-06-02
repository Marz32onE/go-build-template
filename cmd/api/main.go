package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	config "github.com/marz32one/go-build-template/internal"
	api "github.com/marz32one/go-build-template/pkg/api/rest"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	config.Load()

	// Routes
	api.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":9527"))
}
