CREATE TABLE IF NOT EXISTS jogos (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    plataforma TEXT NOT NULL,
    nota NUMERIC(3,1) CHECK (nota >= 0 AND nota <= 10),
    jogado BOOLEAN DEFAULT FALSE
);
