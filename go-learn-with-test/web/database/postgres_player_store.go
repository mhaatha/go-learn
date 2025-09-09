package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPlayerStore() *PostgresPlayerStore {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}

	return &PostgresPlayerStore{
		DB: dbpool,
	}
}

type PostgresPlayerStore struct {
	DB *pgxpool.Pool
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	// Kurang optimal karena melakukan 3 kali pemanggilan ke database.
	// Ada cara yang lebih optimal dengan memanfaatkan fitur yang ada di PostgreSQL
	// if p.GetPlayerScore(name) == 0 {
	// 	_, err := p.DB.Exec(context.Background(), "INSERT INTO users (name, score) VALUES ($1, $2)", name, 1)
	// 	if err != nil {
	// 		fmt.Printf("Eror when RecordWin: %v\n", err)
	// 	}
	// } else {
	// 	_, err := p.DB.Exec(context.Background(), "UPDATE users SET score = score + 1 WHERE name = $1", name)
	// 	if err != nil {
	// 		fmt.Printf("Eror when RecordWin: %v\n", err)
	// 	}
	// }

	_, err := p.DB.Exec(context.Background(), `
	INSERT INTO users (name, score) VALUES ($1, 1)
	ON CONFLICT (name) DO UPDATE SET score = users.score + 1`, name)
	if err != nil {
		fmt.Printf("Eror when RecordWin: %v\n", err)
	}
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	err := p.DB.QueryRow(context.Background(), "SELECT score FROM users WHERE name = $1", name).Scan(&score)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Printf("User not found: %v\n", err)
			return 0
		}
		fmt.Printf("Error when GetPlayerScore: %v\n", err)
		return 0
	}

	return score
}
