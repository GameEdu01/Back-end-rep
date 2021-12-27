package scripts

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PageVariables struct {
	Date string
	Time string
}

// RegisterUserPage is responsible for sending registration page to the front-end
func RegisterUserPage(w http.ResponseWriter, r *http.Request) {
	HomePageVars := PageVariables{}
	t, err := template.ParseFiles("./templates/register.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}

//LoginUserPage is responsible for sending login page to the front end
func LoginUserPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/login.html")
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
		t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		} else {
			username := r.Form["username"]
			password := r.Form["password"]
			fmt.Fprintf(w, "username = %s, password = %s", username, password)
		}
	}
}

//CoursePage is responsible for sending page with course content
func CoursePage(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authToken")
	if !VerifyTokens(authToken) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `incorect token` + `"}`))
		return
	}
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
	content := GetCourseById(DbConnector(), id)

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
func UserCoursesPage(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authToken")
	if !VerifyTokens(authToken) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `incorect token` + `"}`))
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

	content := GetCourseForUser(DbConnector(), id)

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

//MarketPage responsible for giving courses for user to sell
func MarketPage(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authToken")
	if !VerifyTokens(authToken) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `incorect token` + `"}`))
		return
	}
	t, err := template.ParseFiles("./templates/course.html")
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
	content := GetMarketForUser(DbConnector(), id)
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}

//HomePage Todo
func HomePage(w http.ResponseWriter, r *http.Request) {

}

//CoursePost is responsible for getting and saving data about new posts
func CoursePost(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authToken")
	if !VerifyTokens(authToken) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `incorect token` + `"}`))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}

	req := &RequestCourse{}
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
	PostCourse(DbConnector(), req, id)
	return
}

//UserCoursesPost is responsible for getting and saving updates about users courses
func UserCoursesPost(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authToken")
	if !VerifyTokens(authToken) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `incorect token` + `"}`))
		return
	}

	//ToDo
}

//MarketPost ToDo
func MarketPost(w http.ResponseWriter, r *http.Request) {

}

//HomePost ToDo
func HomePost(w http.ResponseWriter, r *http.Request) {

}
