package controllers

import (
	"API/problem2/config"
	"API/problem2/lib/database"
	"API/problem2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all books",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	var books models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	title := c.FormValue("Title")
	publisher := c.FormValue("Publisher")
	year, _ := strconv.Atoi(c.FormValue("Year"))

	var books models.Book
	c.Bind(&books)

	books.Title = title
	books.Publisher = publisher
	books.Year = year

	if err := config.DB.Save(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    books,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	var books models.Book

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	config.DB.Delete(&books, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success detele book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	var books models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	var input models.Book

	title := c.FormValue("Title")
	Publisher := c.FormValue("Publisher")
	year, _ := strconv.Atoi(c.FormValue("Year"))

	input.Title = title
	input.Publisher = Publisher
	input.Year = year

	config.DB.Model(&books).Updates(input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}
