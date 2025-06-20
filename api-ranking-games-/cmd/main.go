package main

import (
	"ranking-games/controller"
	"ranking-games/db"
	"ranking-games/repository"
	"ranking-games/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// camada de banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// camada de repositórios
	GameRepository := repository.NewGameRepository(db)

	// camada de use cases
	GameUseCase := usecase.NewGameUseCase(GameRepository)

	// camada de controlllers
	ProductController := controller.NewGameController(GameUseCase)

	// rotas
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	server.GET("/games", ProductController.GetGames)

	server.POST("/games", ProductController.CreateGame)
	server.GET("/games/:gameId", ProductController.GetGameByID)

	server.PUT("/games/:gameId", ProductController.UpdateGame)

	server.DELETE("/games/:gameId", ProductController.DeleteGame)

	server.Run(":8080")
}
