package controller

import (
	"net/http"
	"strconv"
	"weekly3/entities"
	"weekly3/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type BlogController interface{}

type blogController struct {
	useCase usecase.BlogUsecase
}

func NewBlogController(blogUsecase usecase.BlogUsecase) *blogController {
	return &blogController{
		blogUsecase,
	}
}

func (u *blogController) GetAllBlogs(c echo.Context) error {
	var blogs []entities.Blog
	blog, err := u.useCase.Find(blogs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": blog,
	})
}

func (u *blogController) GetBlogByID(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	data, err := u.useCase.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (u *blogController) CreateBlog(c echo.Context) error {
	var blog entities.Blog
	c.Bind(&blog)
	err := u.useCase.Create(blog)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create new blog",
		"data":    blog,
	})
}

func (u *blogController) UpdateBlog(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	var input entities.Blog
	c.Bind(&input)

	data, err := u.useCase.Update(id, input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "update success!",
		"blog":    data,
	})
}

func (u *blogController) Delete(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	err := u.useCase.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete success!",
	})
}

func (u *blogController) GetBlogByCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("category_id"))

	data, err := u.useCase.FindByCat(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (u *blogController) GetBlogByKey(c echo.Context) error {
	data, err := u.useCase.FindByKey(c.FormValue("keyword"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}
