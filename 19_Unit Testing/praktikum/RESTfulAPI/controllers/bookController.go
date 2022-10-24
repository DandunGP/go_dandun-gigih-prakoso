package controllers

import (
	"net/http"
	"strconv"
	"testing/RESTfulAPI/config"
	"testing/RESTfulAPI/lib/database"
	"testing/RESTfulAPI/models"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all books",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	var books models.Book

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

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
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    books,
	})
}

func DeleteBookController(c echo.Context) error {
	var books models.Book

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&books, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

func UpdateBookController(c echo.Context) error {
	var books models.Book

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Book

	title := c.FormValue("Title")
	Publisher := c.FormValue("Publisher")
	year, _ := strconv.Atoi(c.FormValue("Year"))

	input.Title = title
	input.Publisher = Publisher
	input.Year = year

	if err := config.DB.Model(&books).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}
