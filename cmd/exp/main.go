package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	dbUrl := os.Getenv("myBlogURL")
	if dbUrl == "" {
		fmt.Fprint(os.Stderr, "Unable to get db url from env")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to exec query: %s", err)
	}
	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unbaleto exec query: %s", err)
	}
}
