package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"weekly3/config"
	"weekly3/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type blogSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteBlogs(t *testing.T) {
	suite.Run(t, new(blogSuite))
}

func (s *blogSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	config.DB = gormDB
	s.mock = mock
}

func (s *blogSuite) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

// Valid
func (s *blogSuite) TestBlogController_GetBlog() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       models.Blog
	}{
		{
			name:       "get blog by id",
			path:       "blogs/:id",
			expectCode: http.StatusOK,
			Body: models.Blog{
				Title:   "Cek 1",
				Date:    nil,
				Image:   "Coba.jpg",
				Content: "Ini adalah Content",
				Author: "Dandun",
				UserID: 1,
				CategoryID: 1
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"title", "date", "image", "content", "author", "user_id", "category_id"}).
			AddRow("Cek 1", nil, "Coba.jpg", "Ini adalah content", "Dandun", 1, 1)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE id = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(expectedResult)

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		if s.NoError(GetBlogController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Blog models.Blog `json:"blogs"`
			}

			var Blog Response
			if err := json.Unmarshal(body, &Blog); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].Body, Blog.Blog)
		}
	})
}

func (s *blogSuite) TestBlogController_GetAllBlog() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       []models.Blog
	}{
		{
			name:       "get all Blog",
			path:       "blogs/",
			expectCode: http.StatusOK,
			Body: []models.Blog{{
				Title:   "Cek 1",
				Date:    nil,
				Image:   "Coba.jpg",
				Content: "Ini adalah Content",
				Author: "Dandun",
				UserID: 1,
				CategoryID: 1
			}, {
				Title:   "Cek 2",
				Date:    nil,
				Image:   "Coba2.jpg",
				Content: "Ini adalah Content 2",
				Author: "Dandun2",
				UserID: 2,
				CategoryID: 1
			},
			},
		},
	}
	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"title", "date", "image", "content", "author", "user_id", "category_id"}).
				AddRow("Cek 1", nil, "Coba.jpg", "Ini adalah content", "Dandun", 1, 1).
				AddRow("Cek 2", nil, "Coba2.jpg", "Ini adalah content 2", "Dandun2", 2, 1)
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE `blogs`.`deleted_at` IS NULL")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(GetBlogsController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Blog []models.Blog `json:"blogs"`
				}

				var Blog Response
				if err := json.Unmarshal(body, &Blog); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.Body, Blog.Blog)
			}
		})
	}
}

func (s *blogSuite) TestBlogController_CreateBlog() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.Blog
	}{
		{
			name:       "create Blog",
			path:       "blogs/",
			message:    "success create new Blog",
			expectCode: http.StatusOK,
			Body: models.Blog{
				Title:   "Cek 1",
				Date:    nil,
				Image:   "Coba.jpg",
				Content: "Ini adalah Content",
				Author: "Dandun",
				UserID: 1,
				CategoryID: 1
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `blogs` (`title`, `date`, `image`, `content`, `author`, `user_id`, `category_id`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs("",sqlmock.AnyArg(), "", "", "", 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		if s.NoError(CreateBlogController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string      `json:"message"`
				Blog    models.Blog `json:"blog"`
			}

			var Blog Response
			if err := json.Unmarshal(body, &Blog); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, Blog.Message)
		}
	})
}

func (s *blogSuite) TestBlogController_UpdateBlog() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.Blog
	}{
		{
			name:       "update Blog",
			path:       "blogs/:id",
			message:    "update success",
			expectCode: http.StatusOK,
			Body:       models.Blog{},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `blogs` SET `updated_at`=? WHERE id = ? AND `Blogs`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		if s.NoError(UpdateBlogController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var Blog Response
			if err := json.Unmarshal(body, &Blog); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, Blog.Message)
		}
	})
}

func (s *blogSuite) TestBlogController_DeleteBlog() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
	}{
		{
			name:       "delete Blog",
			path:       "blogs/:id",
			message:    "success delete Blog",
			expectCode: http.StatusOK,
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `Blogs` SET `deleted_at`=? WHERE `Blogs`.`id` = ? AND `Blogs`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath(test[0].path)

		if s.NoError(DeleteBlogController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var Blog Response
			if err := json.Unmarshal(body, &Blog); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, Blog.Message)
		}
	})
}

// Invalid
func (s *blogSuite) TestBlogController_InvalidGetBlog() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get Blog by id error",
			path:       "Blogs/:id",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"name", "email", "password"}).
			AddRow("Dindun", "dandun@gmail.com", "cobatest1")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `Blogs` WHERE id = ? AND `Blogs`.`deleted_at` IS NULL ORDER BY `Blogs`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Record not found!"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := GetBlogController(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *blogSuite) TestBlogController_InvalidGetAllBlog() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get all Blog error",
			path:       "Blogs/",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"name", "email", "password"}).
			AddRow("Dandun", "dandun@gmail.com", "cobatest1").
			AddRow("Didin", "didin@gmail.com", "cobatest2")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `Blogs` WHERE `Blogs`.`deleted_at` IS NULL")).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Record not found!"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)

		err := GetBlogsController(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *blogSuite) TestBlogController_InvalidCreateBlog() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "create Blog error",
			path:           "/Blogs",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `Blogs` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`) VALUES (?,?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "").
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		err := CreateBlogController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *blogSuite) TestBlogController_InvalidUpdateBlog() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "update Blog error",
			path:           "/Blogs/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `Blogs` SET `updated_at`=? WHERE id = ? AND `Blogs`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPut, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := UpdateBlogController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *blogSuite) TestBlogController_InvalidDeleteBlog() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "update Blog error",
			path:           "/Blogs/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `Blogs` SET `deleted_at`=? WHERE `Blogs`.`id` = ? AND `Blogs`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := DeleteBlogController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}
