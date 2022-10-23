package controller

import (
	"belajar-go-echo/config"
	"belajar-go-echo/entities"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type userSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (s *userSuite) SetupSuite() {
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

func (s *userSuite) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

// Vali
func (s *userSuite) TestUserController_GetAllUser() {
	userRepository := repository.NewUserRepository(config.DB)

	userService := usecase.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       []entities.User
	}{
		{
			name:       "get all user",
			path:       "users/",
			expectCode: http.StatusOK,
			Body: []entities.User{{
				Email:    "dandun@gmail.com",
				Password: "cobatest1",
			}, {
				Email:    "didin@gmail.com",
				Password: "cobatest2",
			},
			},
		},
	}
	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"email", "password"}).
				AddRow("dandun@gmail.com", "cobatest1").
				AddRow("didin@gmail.com", "cobatest2")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(userController.GetAllUsers(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					User []entities.User `json:"data"`
				}

				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.Body, user.User)
			}
		})
	}
}

func (s *userSuite) TestUserController_CreateUser() {
	userRepository := repository.NewUserRepository(config.DB)

	userService := usecase.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       entities.User
	}{
		{
			name:       "create user",
			path:       "users/",
			message:    "success create new user",
			expectCode: http.StatusOK,
			Body:       entities.User{},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "").
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		if s.NoError(userController.CreateUser(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string        `json:"message"`
				User    entities.User `json:"data"`
			}

			var user Response
			if err := json.Unmarshal(body, &user); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, user.Message)
		}
	})
}

func (s *userSuite) TestUserController_InvalidGetAllUser() {
	userRepository := repository.NewUserRepository(config.DB)

	userService := usecase.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get all user error",
			path:       "users/",
			expectCode: echo.NewHTTPError(http.StatusInternalServerError, "Error").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"email", "password"}).
			AddRow("dandun@gmail.com", "cobatest1").
			AddRow("didin@gmail.com", "cobatest2")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Error"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)

		err := userController.GetAllUsers(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *userSuite) TestUserController_InvalidCreateUser() {
	userRepository := repository.NewUserRepository(config.DB)

	userService := usecase.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "create user error",
			path:           "/users",
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Error").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "").
			WillReturnError(errors.New("Error"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		err := userController.CreateUser(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}
