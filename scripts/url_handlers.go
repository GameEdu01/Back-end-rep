package scripts

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

type PageVariables struct {
	Date string
	Time string
}

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
