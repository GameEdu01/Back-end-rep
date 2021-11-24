package scripts

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	id       string
	username string
	password string
}

func DbConnector() {
	dbUrl := "postgres://udmehkiskcczbm:d4f6d3d3a48a96f498f7829d75ef285bd9777989c15a135aa5a72903fc86127e@ec2-54-161-164-220.compute-1.amazonaws.com:5432/d2d1ljqhqhl34q"
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	rows, err := db.Query("SELECT username, password FROM logins")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
	users := []User{}

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.username, &user.password); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		users = append(users, user)
	}
	fmt.Printf("found %d user: %+v", len(users), users)
	fmt.Println()
}
