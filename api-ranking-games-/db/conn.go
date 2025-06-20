package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"

	maxRetries    = 10
	retryInterval = 3 * time.Second
)

func ConnectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	for i := 1; i <= maxRetries; i++ {
		log.Printf("Tentativa %d de conexão ao banco de dados...", i)

		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			log.Println("Conexão bem-sucedida com o banco de dados!")
			return db, nil
		}

		log.Printf("Erro ao conectar: %s", err)

		if i < maxRetries {
			log.Printf("Aguardando %s antes de tentar novamente...", retryInterval)
			time.Sleep(retryInterval)
		} else {
			log.Println("Número máximo de tentativas atingido. Falha ao conectar no banco.")
		}
	}

	return nil, fmt.Errorf("falha ao conectar no banco de dados após %d tentativas: %w", maxRetries, err)
}
