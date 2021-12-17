package scripts

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"strconv"
)

type User struct {
	id       string
	username string
	password string
}

type Content struct {
	Request string `json:"request"`
}

type Course struct {
	Id             int     `json:"id"`
	Author_id      int     `json:"author_id"`
	Price          float64 `json:"price"`
	Owners         []int   `json:"owners"`
	Game_name      string  `json:"game_name"`
	Followers      int     `json:"followers"`
	Course_content Content `json:"course_content"`
}

type RequestCourse struct {
	Price          string  `json:"price"`
	Game_name      string  `json:"game_name"`
	Course_content Content `json:"course_content"`
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

func GetLogin(username string, password string) (UserAuth, error) {
	db := DbConnector()
	rows, err := db.Query("SELECT username, password FROM logins WHERE username=$1", username)
	if err != nil {
		return UserAuth{}, err
	}
	var users UserAuth

	for rows.Next() {
		if err := rows.Scan(&users.Username, &users.Password); err != nil {
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

func GetCourseForUser(db *sql.DB, id int) []Course {
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

		fmt.Println(id)
		for i := 0; i < len(course.Owners); i++ {
			if course.Owners[i] == id {
				courses = append(courses, course)
			}
		}

	}
	fmt.Print(courses)
	return courses
}

func PostCourse(db *sql.DB, CoursePosted *RequestCourse, CourseId int) {
	PriceFloat, err := strconv.ParseFloat(CoursePosted.Price, 64)
	if err != nil {
		fmt.Errorf("could not parse price", err)
	}
	var Owners []int
	Owners = append(Owners, CourseId)

	var SaveCourse Course
	SaveCourse.Course_content = CoursePosted.Course_content
	SaveCourse.Game_name = CoursePosted.Game_name
	SaveCourse.Price = PriceFloat
	SaveCourse.Author_id = CourseId
	SaveCourse.Followers = 0
	SaveCourse.Owners = Owners

	result, err := db.Exec("INSERT INTO courses (author_id, price, game_name, followers, course_content, owners) VALUES ($1, $2, $3, $4, $5, $6)", SaveCourse.Author_id, SaveCourse.Price, SaveCourse.Game_name, SaveCourse.Followers, SaveCourse.Course_content, SaveCourse.Owners)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Println(id)
}
