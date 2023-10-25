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
	fmt.Println("Connected to database")

	defer db.Close(ctx)
	_, err = db.Exec(ctx, "delete from orders; ")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(ctx, "select * from users;")
	emailList := []string{}
	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
		}
		emailList = append(emailList, email)
	}
	for id, email := range emailList {
		_, err = db.Exec(ctx, "insert into orders ( user_id, amount, description) values ($1,$2,$3);", id, 100, email)
		if err != nil {
			fmt.Println(err)
		}

	}

}
