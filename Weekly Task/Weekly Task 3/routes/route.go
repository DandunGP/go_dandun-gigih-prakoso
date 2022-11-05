package routes

import (
	"weekly3/constants"
	"weekly3/controller"
	"weekly3/repository"
	"weekly3/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginRepository := repository.NewLoginRepository(db)
	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	blogRepository := repository.NewBlogRepository(db)

	loginService := usecase.NewLoginUsecase(loginRepository)
	userService := usecase.NewUserUsecase(userRepository)
	categoryService := usecase.NewCategoryUsecase(categoryRepository)
	blogService := usecase.NewBlogUsecase(blogRepository)

	loginController := controller.NewLoginController(loginService)
	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	blogController := controller.NewBlogController(blogService)

	e.POST("/login", loginController.GetUser)
	e.POST("/signup", userController.CreateUser)

	eUser := e.Group("users")
	eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	eUser.GET("", userController.GetAllUsers)

	eBlog := e.Group("blogs")
	eBlog.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBlog.GET("", blogController.GetAllBlogs)
	eBlog.GET("/:id", blogController.GetBlogByID)
	eBlog.POST("", blogController.CreateBlog)
	eBlog.PUT("/:id", blogController.UpdateBlog)
	eBlog.DELETE("/:id", blogController.Delete)
	eBlog.GET("/category/:category_id", blogController.GetBlogByCat)
	eBlog.GET("/keyword", blogController.GetBlogByKey)

	eCat := e.Group("category")
	eCat.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	eCat.GET("", categoryController.GetAllCategorys)
	eCat.POST("", categoryController.CreateCategory)
}
