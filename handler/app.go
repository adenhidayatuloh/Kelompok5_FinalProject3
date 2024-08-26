package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/infra/postgres"
	categorypostgres "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/categoryRepository/categoryPostgres"
	taskpostgres "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/taskRepository/taskPostgres"
	userpostgres "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/userrepository/userPostgres"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/service"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func StartApp() {
	db := postgres.GetDBInstance()
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
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

	route.POST("/chatAI", func(c *gin.Context) {

		type GrammarRequest struct {
			Sentence string `json:"sentence"`
		}

		var requestBody GrammarRequest

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			return
		}

		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyC-OrSLtgmj8AZHvBZkD0wQeDulLAAxPfg"))
		if err != nil {
			log.Fatal(err)

			return
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-1.5-flash")
		cs := model.StartChat()

		cs.History = []*genai.Content{
			{
				Parts: []genai.Part{
					genai.Text(`
	Imagine you are a native English speaker. I want to use you as a tool to have a conversation in English as an exercise for students for speaking material. Please answer all user questions below with the following conditions:

	1. The output is a direct answer to the user, an assessment of whether the user's grammar is correct or incorrect, and also how to correct the grammar if the grammar is incorrect.

	2. Each output is made into a json type variable. In the json, there is an "answer" key that will store the output of the answer to the user, the "is_correct" key to store the output of the grammar check, and the "fix" key to store the output of the grammar correction if the grammar is incorrect.

	3. The "is_correct" key contains true or false. If the "is_correct" key = true then the "fix" key is just an empty string and the "answer" key will contain the direct answer to the user. If the "is_correct" key = false then the "fix" key will be filled with a complete correction of what grammar is correct and the "answer" key will contain an empty string.

	4. Ignore grammar checks for capital and punctuation.

	5. No need for json prefix. Go straight to the code

	6. Penggunaan kapital dan tanda baca seperti tanda tanya, koma, titik ataupun seru tolong hiraukan saja pada pernyataan dari user
	`),
				},
				Role: "user",
			},
		}

		res, err := cs.SendMessage(ctx, genai.Text(`Here is the statement:`+requestBody.Sentence))
		if err != nil {
			log.Fatal(err)

			return
		}

		output := printResponse(res)
		var data map[string]interface{}

		// Unmarshal string JSON ke dalam map
		err = json.Unmarshal([]byte(output), &data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.JSON(http.StatusOK, data)
	})

	log.Fatalln(route.Run(":" + port))
}

func printResponse(resp *genai.GenerateContentResponse) string {

	Output := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {

				Output = fmt.Sprint(Output, part)

			}
		}
	}

	Output = strings.Replace(Output, "json", "", -1)
	Output = strings.Replace(Output, "```", "", -1)
	return Output
}
