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
	ctx := context.Background()

	db, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close(ctx)

	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);`,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to exec query: %s", err)
	}

	_, err = db.Exec(ctx, `
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

	var userId int
	err = db.QueryRow(ctx, `insert into users(name, email) values($1, $2) returning id;`, "joy", "joy@mail.com").Scan(&userId)
	if err != nil {
		panic(err)
	}
	fmt.Println(userId)

}
