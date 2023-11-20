package handler

import (
	"finalProject3/infra/postgres"
	categorypostgres "finalProject3/repository/categoryRepository/categoryPostgres"
	taskpostgres "finalProject3/repository/taskRepository/taskPostgres"
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
	categoryRepo := categorypostgres.NewCategoryPG(db)
	taskRepo := taskpostgres.NewTaskPG(db)

	authService := service.NewAuthService(userRepo, taskRepo)

	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	taskService := service.NewTaskService(taskRepo, categoryRepo)
	taskHandler := NewTaskHandler(taskService)

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

	TaskRoute := route.Group("/tasks")
	{
		TaskRoute.POST("/", authService.Authentication(), taskHandler.CreateTask)
		TaskRoute.GET("/", authService.Authentication(), taskHandler.GetAllTasks)
		TaskRoute.PUT("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTask)
		TaskRoute.PATCH("/update-status/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskStatus)
		TaskRoute.PATCH("/update-category/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskCategory)
		TaskRoute.DELETE("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.DeleteTask)
	}

	log.Fatalln(route.Run(":" + port))
}
