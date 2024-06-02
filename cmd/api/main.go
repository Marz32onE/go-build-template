package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	api "github.com/marz32one/go-build-template/pkg/api/rest"
	"github.com/marz32one/go-build-template/pkg/storage"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	storage.InitDB("mydatabase.db")

	// Routes
	api.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":9527"))
}
