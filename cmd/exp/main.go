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
	var firstName string
	var age int
	err = conn.QueryRow(context.Background(), "select first_name, age from users").Scan(&firstName, &age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(firstName, age)
}
