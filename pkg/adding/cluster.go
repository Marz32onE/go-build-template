package adding

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marz32one/go-build-template/pkg/storage"
	data "github.com/marz32one/go-build-template/pkg/storage/data"
)

// CreateItem adds a new item to the database
func CreateItem(c echo.Context) error {
	var item data.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := storage.DB.Create(&item)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, item)
}
