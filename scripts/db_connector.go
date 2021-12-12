package scripts

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"strings"
)

type User struct {
	id       string
	username string
	password string
}

type Course struct {
	Id             string `json:"id"`
	Author_id      string `json:"author_id"`
	Price          string `json:"price"`
	Owners         string `json:"owners"`
	Game_name      string `json:"game_name"`
	Followers      string `json:"followers"`
	Course_content string `json:"course_content"`
}

func DbConnector() *sql.DB {
	dbUrl := "postgres://udmehkiskcczbm:d4f6d3d3a48a96f498f7829d75ef285bd9777989c15a135aa5a72903fc86127e@ec2-54-161-164-220.compute-1.amazonaws.com:5432/d2d1ljqhqhl34q"
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")
	return db
}

func GetLogin(db *sql.DB, username string, password string) []User {
	rows, err := db.Query("SELECT username, password FROM logins WHERE username=$1 and password=$2", username, password)
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
	var users []User

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.username, &user.password); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		users = append(users, user)
	}
	fmt.Printf("found %d user: %+v", len(users), users)
	fmt.Println()
	return users

}

func CreateUserInDB(db *sql.DB, username string, password string) {
	//ToDo
}

func GetCourseById(db *sql.DB, id uuid.UUID) Course {
	row := db.QueryRow("SELECT * FROM courses WHERE id=$1", id)
	course := Course{}
	if err := row.Scan(
		&course.Id, &course.Author_id,
		&course.Price, &course.Game_name,
		&course.Followers, &course.Course_content,
		&course.Owners,
	); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}

	fmt.Printf("found course:", course)
	fmt.Println()
	return course
}

func GetCourseForUser(db *sql.DB, id uuid.UUID) []Course {
	rows, err := db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	var courses []Course

	for rows.Next() {
		course := Course{}
		if err := rows.Scan(
			&course.Id, &course.Author_id,
			&course.Price, &course.Game_name,
			&course.Followers, &course.Course_content,
			&course.Owners,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

		fmt.Println(id.String())
		if strings.Contains(course.Owners, id.String()) {
			courses = append(courses, course)
		}
	}
	fmt.Print(courses)
	return courses
}
