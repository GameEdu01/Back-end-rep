package db

import (
	"database/sql"
	Types "eduapp/CommonTypes"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

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

func GetLogin(username string) (Types.User, error) {
	db := DbConnector()
	rows, err := db.Query("SELECT * FROM logins WHERE username=$1", username)
	if err != nil {
		return Types.User{}, err
	}
	var users Types.User

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Username, &users.Password); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

	}
	fmt.Printf("found %d user: %+v", users)
	fmt.Println()
	return users, nil

}

func CreateUserInDB(username string, password string) {
	db := DbConnector()
	result, err := db.Exec("INSERT INTO logins (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Println(id)
}

func GetCourseById(db *sql.DB, id uuid.UUID) Types.Course {
	row := db.QueryRow("SELECT * FROM courses WHERE id=$1", id)
	course := Types.Course{}
	if err := row.Scan(
		&course.Id, &course.Author_id,
		&course.Category, &course.Game,
		&course.Description, &course.Image,
		&course.Language, &course.PublishedAt,
		&course.Title, &course.Url,
	); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}

	fmt.Printf("found course:", course)
	fmt.Println()
	return course
}

func GetNewsFeed(db *sql.DB, ammount int) []Types.Course {
	rows, err := db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	var courses []Types.Course

	for rows.Next() {
		course := Types.Course{}
		if err := rows.Scan(
			&course.Id, &course.Author_id,
			&course.Category, &course.Game,
			&course.Description, &course.Image,
			&course.Language, &course.PublishedAt,
			&course.Title, &course.Url,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		fmt.Println(ammount)
		for i := 0; i < ammount; i++ {
			courses = append(courses, course)
		}
	}
	fmt.Println(courses)
	return courses
}

func GetCourseForUser(db *sql.DB, id int) []Types.Course {
	rows, err := db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	var courses []Types.Course

	for rows.Next() {
		course := Types.Course{}
		if err := rows.Scan(
			&course.Id, &course.Author_id,
			&course.Category, &course.Game,
			&course.Description, &course.Image,
			&course.Language, &course.PublishedAt,
			&course.Title, &course.Url,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
	}
	fmt.Print(courses)
	return courses
}

func PostCourse(db *sql.DB, CoursePosted *Types.Course) {
	var SaveCourse Types.Course
	SaveCourse.Author_id = CoursePosted.Author_id
	SaveCourse.Category = CoursePosted.Category
	SaveCourse.Game = CoursePosted.Game
	SaveCourse.Description = CoursePosted.Description
	SaveCourse.Image = CoursePosted.Image
	SaveCourse.Language = CoursePosted.Language
	SaveCourse.PublishedAt = CoursePosted.PublishedAt
	SaveCourse.Title = CoursePosted.Title
	SaveCourse.Url = CoursePosted.Url

	result, err := db.Exec("INSERT INTO courses (author_id, category_game, game, description, image, language_content, published_at, title, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", SaveCourse.Author_id, SaveCourse.Category, SaveCourse.Game, SaveCourse.Description, SaveCourse.Image, SaveCourse.Language, SaveCourse.PublishedAt, SaveCourse.Title, SaveCourse.Url)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Println(id)
}

func GetMarketForUser(db *sql.DB, id int) []Types.Course {
	rows, err := db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	var courses []Types.Course

	for rows.Next() {
		course := Types.Course{}
		if err := rows.Scan(
			&course.Id, &course.Author_id,
			&course.Category, &course.Game,
			&course.Description, &course.Image,
			&course.Language, &course.PublishedAt,
			&course.Title, &course.Url,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

		fmt.Println(id)

		//ToDo selecting courses for user

	}
	fmt.Print(courses)
	return courses
}
