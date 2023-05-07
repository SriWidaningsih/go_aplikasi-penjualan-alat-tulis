package controllers

import (
	"miniprojek/models"
	"net/http"

	"miniprojek/config"

	"github.com/labstack/echo/v4"
)

// get all products
func GetProductsController(c echo.Context) error {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all products",
		"users":   products,
	})
}
