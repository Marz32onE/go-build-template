package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marz32one/go-build-template/pkg/storage"
	"github.com/marz32one/go-build-template/pkg/storage/data"
	"gorm.io/gorm"
)

// CreateItem adds a new item to the database
// @Summary Add a new item
// @Description Add a new item to the database
// @Tags Items
// @Accept json
// @Produce json
// @Param item body data.Resource true "Item object to be added"
// @Success 201 {object} data.Resource
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items [post]
func CreateItem(c echo.Context) error {
	var item data.Resource
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := storage.DB.Create(&item)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, item)
}

// GetItems retrieves all items from the database
// @Summary Get all items
// @Description Retrieve all items from the database
// @Tags Items
// @Produce json
// @Success 200 {array} data.Resource
// @Failure 500 {object} string
// @Router /items [get]
func GetItems(c echo.Context) error {
	var items []data.Resource
	result := storage.DB.Find(&items)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, items)
}

// GetItem retrieves a single item by ID from the database
// @Summary Get an item by ID
// @Description Retrieve a single item by ID from the database
// @Tags Items
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} data.Resource
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [get]
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
