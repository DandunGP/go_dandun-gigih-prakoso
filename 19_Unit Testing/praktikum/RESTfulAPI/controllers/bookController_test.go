package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"testing/RESTfulAPI/config"
	"testing/RESTfulAPI/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type bookSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteBooks(t *testing.T) {
	suite.Run(t, new(bookSuite))
}

func (s *bookSuite) SetupSuite() {
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

func (s *bookSuite) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

// Valid
func (s *bookSuite) TestBookController_GetBook() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       models.Book
	}{
		{
			name:       "get book by id",
			path:       "books/:id",
			expectCode: http.StatusOK,
			Body: models.Book{
				Title:     "Malin",
				Publisher: "Sidu",
				Year:      2022,
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"title", "publisher", "year"}).
			AddRow("Malin", "Sidu", 2022)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE id = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
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

		if s.NoError(GetBookController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Book models.Book `json:"books"`
			}

			var book Response
			if err := json.Unmarshal(body, &book); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].Body, book.Book)
		}
	})
}

func (s *bookSuite) TestBookController_GetAllBook() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       []models.Book
	}{
		{
			name:       "get all book",
			path:       "books/",
			expectCode: http.StatusOK,
			Body: []models.Book{{
				Title:     "Malin",
				Publisher: "Sidu",
				Year:      2022,
			}, {
				Title:     "Mulan",
				Publisher: "Sudi",
				Year:      2001,
			},
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"title", "publisher", "year"}).
			AddRow("Malin", "Sidu", 2022).
			AddRow("Mulan", "Sudi", 2001)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL")).
			WillReturnRows(expectedResult)

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)

		if s.NoError(GetBooksController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Book []models.Book `json:"books"`
			}

			var book Response
			if err := json.Unmarshal(body, &book); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].Body, book.Book)
		}
	})
}

func (s *bookSuite) TestBookController_CreateBook() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.Book
	}{
		{
			name:       "create book",
			path:       "books/",
			message:    "success create new book",
			expectCode: http.StatusOK,
			Body: models.Book{
				Title:     "Malin",
				Publisher: "Sidu",
				Year:      2022,
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`title`,`publisher`,`year`) VALUES (?,?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", 0).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		if s.NoError(CreateBookController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string      `json:"message"`
				Book    models.Book `json:"book"`
			}

			var book Response
			if err := json.Unmarshal(body, &book); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, book.Message)
		}
	})
}

func (s *bookSuite) TestBookController_UpdateBook() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.Book
	}{
		{
			name:       "update user",
			path:       "books/:id",
			message:    "update success",
			expectCode: http.StatusOK,
			Body:       models.Book{},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=? WHERE id = ? AND `books`.`deleted_at` IS NULL")).
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

		if s.NoError(UpdateBookController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var book Response
			if err := json.Unmarshal(body, &book); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, book.Message)
		}
	})
}

func (s *bookSuite) TestBookController_DeleteBook() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
	}{
		{
			name:       "delete book",
			path:       "books/:id",
			message:    "success delete book",
			expectCode: http.StatusOK,
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `deleted_at`=? WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL")).
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

		if s.NoError(DeleteBookController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var book Response
			if err := json.Unmarshal(body, &book); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, book.Message)
		}
	})
}

// Invalid
func (s *bookSuite) TestBookController_InvalidGetBook() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get book by id error",
			path:       "books/:id",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"title", "publisher", "year"}).
			AddRow("Malin", "Sidu", 2022)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE id = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
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

		err := GetBookController(ctx).Error()

		fmt.Println(err)
		s.Equal(test[0].expectCode, err)
	})
}

func (s *bookSuite) TestBookController_InvalidGetAllBook() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get all book error",
			path:       "books/",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"title", "publisher", "year"}).
			AddRow("Malin", "Sidu", 2022).
			AddRow("Mulin", "Sudi", 2001)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL")).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Record not found!"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)

		err := GetBooksController(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *bookSuite) TestBookController_InvalidCreateBook() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "create book error",
			path:           "/books",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`title`,`publisher`,`year`) VALUES (?,?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", 0).
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		err := CreateBookController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *bookSuite) TestBookController_InvalidUpdateBook() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "update book error",
			path:           "/books/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=? WHERE id = ? AND `books`.`deleted_at` IS NULL")).
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

		err := UpdateBookController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *bookSuite) TestBookController_InvalidDeleteBook() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "delete user error",
			path:           "/books/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `deleted_at`=? WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL")).
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

		err := DeleteBookController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}
