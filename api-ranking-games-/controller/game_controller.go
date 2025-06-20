package controller

import (
	"net/http"
	"ranking-games/model"
	"ranking-games/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type gameController struct {
	//adicionar dependências, como use cases ou serviços
	gameUseCase usecase.GameUseCase
}

func NewGameController(useCase usecase.GameUseCase) gameController {
	return gameController{
		gameUseCase: useCase,
	}
}

func (gc *gameController) GetGames(ctx *gin.Context) {

	games, err := gc.gameUseCase.GetGames()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving games",
			"status":  "error",
		})
		return
	}

	if len(games) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No games found",
			"status":  "not_found",
		})
		return
	}

	//se tudo estiver certo, retorna os jogos

	ctx.JSON(http.StatusOK, gin.H{
		"games":   games,
		"message": "Games retrieved successfully",
		"status":  "success",
	})
}

func (gc *gameController) CreateGame(ctx *gin.Context) {
	var gameInput model.Game

	if err := ctx.ShouldBindJSON(&gameInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"status":  "error",
		})
		return
	}

	game, err := gc.gameUseCase.CreateGame(gameInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating game",
			"status":  "error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"game":    game,
		"message": "Game created successfully",
		"status":  "success",
	})
}

func (gc *gameController) GetGameByID(ctx *gin.Context) {
	id, exists := ctx.Params.Get("gameId")

	if !exists || id == "" {
		response := model.Response{
			Message: "Game ID is required",
			Status:  "error",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	gameId, convErr := strconv.Atoi(id)

	if convErr != nil {
		response := model.Response{
			Message: "Invalid Game ID format",
			Status:  "error",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	game, err := gc.gameUseCase.GetGameByID(gameId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving game",
			"status":  "error",
		})
		return
	}

	if game == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Game not found",
			"status":  "not_found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    game,
		"message": "Game retrieved successfully",
		"status":  "success",
	})
}

func (gc *gameController) UpdateGame(ctx *gin.Context) {
	id := ctx.Param("gameId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Game ID is required",
			"status":  "error",
		})
		return
	}

	gameId, convErr := strconv.Atoi(id)
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Game ID format",
			"status":  "error",
		})
		return
	}

	var gameInput model.Game
	if err := ctx.ShouldBindJSON(&gameInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"status":  "error",
		})
		return
	}

	gameInput.ID = gameId

	updatedGame, err := gc.gameUseCase.UpdateGame(gameInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating game: " + err.Error(),
			"status":  "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"game":    updatedGame,
		"message": "Game updated successfully",
		"status":  "success",
	})
}

func (gc *gameController) DeleteGame(ctx *gin.Context) {
	id := ctx.Param("gameId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Game ID is required",
			"status":  "error",
		})
		return
	}

	gameId, convErr := strconv.Atoi(id)
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Game ID format",
			"status":  "error",
		})
		return
	}

	err := gc.gameUseCase.DeleteGame(gameId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting game: " + err.Error(),
			"status":  "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Game deleted successfully",
		"status":  "success",
	})
}
