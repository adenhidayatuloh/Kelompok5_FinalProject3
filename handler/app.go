package handler

import (
	"finalProject3/infra/postgres"
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

	UsersRoute := route.Group("/users")
	{
		UsersRoute.POST("/register", userHandler.Register)
		UsersRoute.POST("/login", userHandler.Login)
		UsersRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)
		UsersRoute.DELETE("/delete-account", authService.Authentication(), userHandler.DeleteUser)
	}

	log.Fatalln(route.Run(":" + port))
}
