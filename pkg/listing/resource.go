package listing

import (
	"net/http"

	"github.com/marz32one/go-build-template/pkg/storage"
	data "github.com/marz32one/go-build-template/pkg/storage/data"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetItems retrieves all items from the database
func GetItems(c echo.Context) error {
	var items []data.Resource
	result := storage.DB.Find(&items)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, items)
}

// GetItem retrieves a single item by ID from the database
func GetItem(c echo.Context) error {
	id := c.Param("id")
	var item data.Resource
	result := storage.DB.First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Item not found")
		}
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, item)
}
