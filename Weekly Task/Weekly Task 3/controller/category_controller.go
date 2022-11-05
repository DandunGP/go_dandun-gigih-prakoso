package controller

import (
	"net/http"
	"weekly3/entities"
	"weekly3/usecase"

	"github.com/labstack/echo"
)

type CategoryController interface{}

type categoryController struct {
	useCase usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *categoryController {
	return &categoryController{
		categoryUsecase,
	}
}

func (u *categoryController) GetAllCategorys(c echo.Context) error {
	var categorys []entities.Category
	category, err := u.useCase.Find(categorys)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": category,
	})
}

func (u *categoryController) CreateCategory(c echo.Context) error {
	var category entities.Category
	c.Bind(&category)
	err := u.useCase.Create(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create new Category",
		"data":    category,
	})
}
