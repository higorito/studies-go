package repository

import (
	"database/sql"
	"fmt"
	"ranking-games/model"
)

type GameRepository struct {
	connection *sql.DB
}

func NewGameRepository(db *sql.DB) GameRepository {
	return GameRepository{
		connection: db,
	}
}

func (gr *GameRepository) GetGames() ([]model.Game, error) {
	query := "SELECT id, nome, plataforma, nota, jogado FROM jogos"

	rows, err := gr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Game{}, err
	}

	var games []model.Game
	var gameObj model.Game

	for rows.Next() {
		err := rows.Scan(&gameObj.ID, &gameObj.Nome, &gameObj.Plataforma, &gameObj.Nota, &gameObj.Jogado)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Game{}, err
		}

		games = append(games, gameObj)
	}

	rows.Close()

	return games, nil
}

func (gr *GameRepository) CreateGame(game model.Game) (int, error) {
	query, err := gr.connection.Prepare("INSERT INTO jogos" +
		"(nome, plataforma, nota, jogado) VALUES ($1, $2, $3, $4) RETURNING id")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}

	var gameID int

	err = query.QueryRow(game.Nome, game.Plataforma, game.Nota, game.Jogado).Scan(&gameID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}

	return gameID, nil
}

func (gr *GameRepository) GetGameByID(id int) (*model.Game, error) {
	query, err := gr.connection.Prepare("SELECT id, nome, plataforma, nota, jogado FROM jogos WHERE id = $1")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return nil, err
	}

	var gameObj model.Game

	err = query.QueryRow(id).Scan(&gameObj.ID, &gameObj.Nome, &gameObj.Plataforma, &gameObj.Nota, &gameObj.Jogado)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	return &gameObj, nil
}

func (gr *GameRepository) UpdateGame(game model.Game) error {
	query, err := gr.connection.Prepare(`
		UPDATE jogos 
		SET nome = $1, plataforma = $2, nota = $3, jogado = $4 
		WHERE id = $5
	`)
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return fmt.Errorf("error preparing update query: %w", err)
	}
	defer query.Close()

	result, err := query.Exec(game.Nome, game.Plataforma, game.Nota, game.Jogado, game.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return fmt.Errorf("error executing update: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return fmt.Errorf("no game found with ID %d", game.ID)
	}

	return err
}

func (gr *GameRepository) DeleteGame(id int) error {
	query, err := gr.connection.Prepare("DELETE FROM jogos WHERE id = $1")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return fmt.Errorf("error preparing delete query: %w", err)
	}
	defer query.Close()

	result, err := query.Exec(id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return fmt.Errorf("error executing delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return fmt.Errorf("no game found with ID %d", id)
	}

	return err
}
