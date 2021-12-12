package main

import (
	"eduapp/scripts"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./templates/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")

	// Routers GET
	r.HandleFunc("/register", scripts.RegisterUserPage).Methods("GET")
	r.HandleFunc("/login", scripts.LoginUserPage).Methods("GET")
	r.HandleFunc("/course", scripts.CoursePage).Methods("GET")
	r.HandleFunc("/mycourses", scripts.UserCoursesPage).Methods("GET")
	r.HandleFunc("/market", scripts.MarketPage).Methods("GET")
	r.HandleFunc("/home", scripts.HomePage).Methods("GET")

	//Routers POST
	r.HandleFunc("/api/user/login", scripts.UserLogin).Methods("POST")
	r.HandleFunc("/api/user/signup", scripts.UserSignup).Methods("POST")
	r.HandleFunc("/course", scripts.CoursePost).Methods("POST")
	r.HandleFunc("/mycourses", scripts.UserCoursesPost).Methods("POST")
	r.HandleFunc("/market", scripts.MarketPost).Methods("POST")
	r.HandleFunc("/home", scripts.HomePost).Methods("POST")

	return r
}

func main() {
	port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	scripts.DbConnector()
	r := newRouter()
	fmt.Println("hello i am started")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
