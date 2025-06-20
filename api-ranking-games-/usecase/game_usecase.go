package usecase

import (
	"ranking-games/model"
	"ranking-games/repository"
)

// se for maiuscula é exportado, se for minuscula é privado(dentro do pacote)
type GameUseCase struct {
	repository repository.GameRepository
}

func NewGameUseCase(repo repository.GameRepository) GameUseCase {
	return GameUseCase{
		repository: repo,
	}
}

func (gu *GameUseCase) GetGames() ([]model.Game, error) {

	return gu.repository.GetGames()
}

func (gu *GameUseCase) CreateGame(game model.Game) (model.Game, error) {
	gameID, err := gu.repository.CreateGame(game)
	if err != nil {
		return model.Game{}, err
	}

	game.ID = gameID
	return game, nil
}

func (gu *GameUseCase) GetGameByID(id int) (*model.Game, error) {
	game, err := gu.repository.GetGameByID(id)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (gu *GameUseCase) UpdateGame(game model.Game) (model.Game, error) {
	err := gu.repository.UpdateGame(game)
	if err != nil {
		return model.Game{}, err
	}

	return game, nil
}

func (gu *GameUseCase) DeleteGame(id int) error {
	err := gu.repository.DeleteGame(id)
	if err != nil {
		return err
	}

	return nil
}
