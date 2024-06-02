package api

import (
	"github.com/labstack/echo/v4"
	"github.com/marz32one/go-build-template/pkg/adding"
	"github.com/marz32one/go-build-template/pkg/listing"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(e *echo.Echo) {
	e.GET("/items", listing.GetItems)
	e.GET("/items/:id", listing.GetItem)
	e.POST("/items", adding.CreateItem)
}
