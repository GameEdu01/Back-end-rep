package scripts

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
	t, err := template.ParseFiles("./templates/Login.html")
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

//MarketPage ToDo
func MarketPage(w http.ResponseWriter, r *http.Request) {

}

//HomePage Todo
func HomePage(w http.ResponseWriter, r *http.Request) {

}

//CoursePost is responsible for getting and saving data about new posts
func CoursePost(w http.ResponseWriter, r *http.Request) {
	request := Course{
		Id:             "id",
		Author_id:      "author_id",
		Price:          "price",
		Owners:         "owners",
		Game_name:      "game_name",
		Followers:      "followers",
		Course_content: "course_content",
	}

	fmt.Println(request)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	req := &Course{}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}
	fmt.Printf("%+v\n", req)

	//ToDo save to db

}

//UserCoursesPost is responsible for getting and saving updates about users courses
func UserCoursesPost(w http.ResponseWriter, r *http.Request) {

}

//MarketPost ToDo
func MarketPost(w http.ResponseWriter, r *http.Request) {

}

//HomePost ToDo
func HomePost(w http.ResponseWriter, r *http.Request) {

}
