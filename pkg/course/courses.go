package course

import (
	Types "eduapp/CommonTypes"
	"eduapp/pkg/db"
	myerrors "eduapp/pkg/errors"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//CoursePage is responsible for sending page with course content
func CoursePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	id, err := uuid.Parse(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	content := db.GetCourseById(db.DbConnector(), id)

	t, err := template.ParseFiles("./templates/Course.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SendNewsFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Feed := db.GetNewsFeed(db.DbConnector(), 1)
	b, err := json.Marshal(Feed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `response parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	fmt.Println(string(b))
	w.Write(b)
}

func NewsFeedPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/NewsFeed.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
}

//UserCoursesPage is responsible for sending courses owned by user
func UserCoursesPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	id, err := strconv.Atoi(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	content := db.GetCourseForUser(db.DbConnector(), id)

	t, err := template.ParseFiles("./templates/Course.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}

//PostCourse is responsible for getting and saving data about new posts
func PostCourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	req := &Types.Course{}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	req.Language = "en"
	req.PublishedAt = time.Now().String()
	req.Views = 1
	fmt.Printf("%+v\n", req)
	db.PostCourse(db.DbConnector(), req)
	return
}

// PagePostCourse is responsible for getting and saving data about new posts
func PagePostCourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/PostCourse.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
}
