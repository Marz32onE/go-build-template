package router

import (
	"github.com/labstack/echo/v4"
	_ "github.com/marz32one/go-build-template/docs" // This imports the generated documentation.
	"github.com/marz32one/go-build-template/pkg/api/handler"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// InitRoutes initializes the routes for the API.
func InitRoutes(e *echo.Echo) {
	apiv1 = e.Group("/api/v1")

	apiv1.GET("/items", handler.GetItems)
	apiv1.GET("/items/:id", handler.GetItem)
	apiv1.POST("/items", handler.CreateItem)

	apiv1.GET("/nodes", handler.GetNodeResources)

	apiv1.GET("/swagger/*", echoSwagger.WrapHandler)
}
