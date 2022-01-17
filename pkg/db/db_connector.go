package db

import (
	"context"
	"database/sql"
	Types "eduapp/CommonTypes"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"time"
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

func GetCourseById(db *sql.DB, id int) Types.Course {
	row := db.QueryRow("SELECT * FROM courses WHERE id=$1", id)
	course := Types.Course{}

	if err := row.Scan(
		&course.Id, &course.Author_id,
		&course.Category, &course.Game,
		&course.Description, &course.Image,
		&course.Language, &course.PublishedAt,
		&course.Title, &course.Content,
		&course.Views,
	); err != nil {
		fmt.Println("could not scan row: %v", err)
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
			&course.Title, &course.Content, &course.Views,
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
			&course.Title, &course.Content,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
	}
	fmt.Print(courses)
	return courses
}

func PostCourse(db *sql.DB, CoursePosted *Types.Course) int64 {
	query := "INSERT INTO courses (author_id, category_game, game, description, image, language_content, published_at, title, content, views) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, CoursePosted.Author_id, CoursePosted.Category, CoursePosted.Game, CoursePosted.Description, CoursePosted.Image, CoursePosted.Language, CoursePosted.PublishedAt, CoursePosted.Title, CoursePosted.Content, CoursePosted.Views)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return 0
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0
	}
	log.Printf("%d products created ", rows)

	courseId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted product", err)
		return 0
	}
	log.Printf("Product with ID %d created", courseId)
	return courseId

}

func PostContent(db *sql.DB, id int, content string) {
	sqlStatement := `
UPDATE courses
SET content = $2
WHERE id = $1;`

	err := db.QueryRow(sqlStatement, id, content)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
}
