package handler

import (
	Types "eduapp/CommonTypes"
	"eduapp/pkg/db"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//CoursePage is responsible for sending page with course content
func CoursePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		return
	}

	id, err := uuid.Parse(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		return
	}
	content := db.GetCourseById(db.DbConnector(), id)

	t, err := template.ParseFiles("./templates/course.html")
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}

//UserCoursesPage is responsible for sending
func UserCoursesPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	content := db.GetCourseForUser(db.DbConnector(), id)

	t, err := template.ParseFiles("./templates/course.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}

//CoursePost is responsible for getting and saving data about new posts
func CoursePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}

	req := &Types.RequestCourse{}
	err = json.Unmarshal(body, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}
	fmt.Printf("%+v\n", req)
	id, err := GetIdByLogin(r.Header.Get("username"))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message":"` + `unable to find user in db` + `"}`))
	}
	fmt.Print(id)
	db.PostCourse(db.DbConnector(), req, id)
	return
}

//UserCoursesPost is responsible for getting and saving updates about users courses
func UserCoursesPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//ToDo
}
