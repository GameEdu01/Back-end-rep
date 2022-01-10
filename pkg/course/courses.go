package course

import (
	Types "eduapp/CommonTypes"
	"eduapp/pkg/db"
	myerrors "eduapp/pkg/errors"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//PageCourse is responsible for sending page with course content
func PageCourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	id, _ := strconv.Atoi(string(b))
	id = 13
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	content := db.GetCourseById(db.DbConnector(), id)
	t, err := template.ParseFiles("./templates/Course.html")
	fmt.Println(err)
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

	resivedCourse := &Types.ResivedCourse{}
	err = json.Unmarshal(body, resivedCourse)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	req := &Types.Course{}
	req.Title = resivedCourse.Title
	req.Description = resivedCourse.Description
	req.Game = resivedCourse.Game
	req.Category = resivedCourse.Category
	req.Image = resivedCourse.Image
	req.Language = "en"
	req.PublishedAt = time.Now().String()
	req.Views = 1

	fmt.Printf("%+v\n", req)

	id := db.PostCourse(db.DbConnector(), req)
	fmt.Printf("%+v\n", id)
	w.Write([]byte(`{"id":"` + strconv.Itoa(0) + `"}`))
	w.WriteHeader(http.StatusOK)
}

func PostContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	resivedContent := Types.Content{}
	err = json.Unmarshal(body, &resivedContent)
	fmt.Printf("%+v\n", resivedContent)
	//ToDo: save content to db
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"` + `susses` + `"}`))

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

func PagePostContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/CreateContent.html")
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
