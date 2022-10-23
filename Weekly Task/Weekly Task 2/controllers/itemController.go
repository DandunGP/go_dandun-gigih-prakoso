package controllers

import (
	"net/http"
	"weekly2/config"
	"weekly2/models"

	"github.com/labstack/echo"
)

func GetItemsController(c echo.Context) error {
	var items []models.Item

	if err := config.DB.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"items":   items,
	})
}

func GetItemController(c echo.Context) error {
	var items models.Item

	if err := config.DB.Where("id = ?", c.Param("id")).First(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

func CreateItemController(c echo.Context) error {
	var items models.Item
	c.Bind(&items)

	if err := config.DB.Save(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    items,
	})
}

func UpdateItemController(c echo.Context) error {
	var items models.Item

	if err := config.DB.Where("id = ?", c.Param("id")).First(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	var input models.Item
	c.Bind(&input)

	config.DB.Model(&items).Update(input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteItemController(c echo.Context) error {
	var items models.Item

	if err := config.DB.Where("id = ?", c.Param("id")).First(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	config.DB.Delete(&items, c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success detele Item",
	})
}

func GetItemByCategory(c echo.Context) error {
	var items []models.Item

	if err := config.DB.Where("category = ?", c.Param("category_id")).Find(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"category": c.Param("category_id"),
		"items":    items,
	})
}

func GetItemByName(c echo.Context) error {
	var items []models.Item

	keyword := c.FormValue("keyword")

	if err := config.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"keyword": keyword,
		"items":   items,
	})
}
