package model

// CREATE TABLE jogos (
//     id SERIAL PRIMARY KEY,
//     nome TEXT NOT NULL,
//     plataforma TEXT NOT NULL,
//     nota NUMERIC(3,1) CHECK (nota >= 0 AND nota <= 10),
//     jogado BOOLEAN DEFAULT FALSE
// );

type Game struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Plataforma string  `json:"plataforma"`
	Nota       float64 `json:"nota"`
	Jogado     bool    `json:"jogado"`
}
