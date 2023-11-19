package handler

import (
	"finalProject3/infra/postgres"
	categorypostgres "finalProject3/repository/categoryRepository/categoryPostgres"
	userpostgres "finalProject3/repository/userrepository/userPostgres"
	"finalProject3/service"
	"log"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := postgres.GetDBInstance()

	port := "8080"
	route := gin.Default()

	userRepo := userpostgres.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)
	authService := service.NewAuthService(userRepo)

	categoryRepo := categorypostgres.NewCategoryPG(db)

	categoryService := service.NewCategoryService(categoryRepo)

	categoryHandler := NewCategoryHandler(categoryService)

	UsersRoute := route.Group("/users")
	{
		UsersRoute.POST("/register", userHandler.Register)
		UsersRoute.POST("/login", userHandler.Login)
		UsersRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)
		UsersRoute.DELETE("/delete-account", authService.Authentication(), userHandler.DeleteUser)
	}

	CategoryRoute := route.Group("/categories")
	{
		CategoryRoute.POST("/", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.CreateCategory)
		CategoryRoute.GET("/", categoryHandler.GetAllCategories)
		CategoryRoute.PATCH("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.UpdateCategory)
		CategoryRoute.DELETE("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.DeleteCategory)
	}

	log.Fatalln(route.Run(":" + port))
}
