package scripts

import (
	"fmt"
	uuid "github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"
)

type login struct {
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
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//LoginUserPage is responsible for sending login page to the front end
func LoginUserPage(w http.ResponseWriter, r *http.Request) {
	HomePageVars := PageVariables{}
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//CoursePage is responsible for sending page with course content
func CoursePage(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("request parsing error: ", err)
	}

	id, err := uuid.Parse(string(b))
	if err != nil {
		log.Print("error parsing UUID ", err)
	}
	content := GetCourseById(DbConnector(), id)

	t, err := template.ParseFiles("./templates/course.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, content)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//UserCoursesPage is responsible for sending
func UserCoursesPage(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("request parsing error: ", err)
	}

	id, err := uuid.Parse(string(b))
	if err != nil {
		log.Print("error parsing UUID ", err)
	}
	content := GetCourseForUser(DbConnector(), id)

	t, err := template.ParseFiles("./templates/course.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, content)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func MarketPage(w http.ResponseWriter, r *http.Request) {

}

func HomePage(w http.ResponseWriter, r *http.Request) {

}

// RegisterUser is responsible for getting user data from the front-end and saving to DB
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := login{
		Username: "username",
		Password: "password",
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request.Username = r.Form.Get("username")
	request.Password = r.Form.Get("password")

	//Todo add user to db

	fmt.Println(request.Username)
	fmt.Println(request.Password)
}

// LoginUser is responsible for getting user data and changing state to loggined
func LoginUser(w http.ResponseWriter, r *http.Request) {
	request := login{
		Username: "username",
		Password: "password",
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request.Username = r.Form.Get("username")
	request.Password = r.Form.Get("password")

	//Todo find user in db

	fmt.Println(request.Username)
	fmt.Println(request.Password)
}

func CoursePost(w http.ResponseWriter, r *http.Request) {
}

func UserCoursesPost(w http.ResponseWriter, r *http.Request) {

}

func MarketPost(w http.ResponseWriter, r *http.Request) {

}

func HomePost(w http.ResponseWriter, r *http.Request) {

}
