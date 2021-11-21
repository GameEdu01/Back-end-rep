package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	fmt.Println("router created")
	r.HandleFunc("/hello", handler).Methods("GET")
	staticFileDirectory := http.Dir("./templates/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/request", getHandler).Methods("GET")
	r.HandleFunc("/request", createHandler).Methods("POST")
	return r
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := newRouter()
	fmt.Println("hello i am started")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
