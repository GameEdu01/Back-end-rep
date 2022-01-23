package future

import (
	"database/sql"
	Types "eduapp/CommonTypes"
	"eduapp/pkg/db"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

//MarketPage responsible for giving courses for user to sell
//Future functional
func MarketPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := template.ParseFiles("./templates/Course.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		return
	}
	id, err := strconv.Atoi(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		return
	}
	content := GetMarketForUser(db.DbConnector(), id)
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}

//Leaderboard
//Future functional
func Leaderboard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/nav.html")
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
		err = t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
	}
}

func NewMain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/profileRedesign.html")
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
		err = t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
	}
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
			&course.Title, &course.Content,
		); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

		fmt.Println(id)

		//selecting courses for user

	}
	fmt.Print(courses)
	return courses
}
