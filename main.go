package main

import (
	"fmt"
	scripts "gited/EduApp/scripts"
	"html/template"
	"net/http"
	"time"
)

// RequestFromFront create a struct that holds information from request
type RequestFromFront struct {
	DataRequested string
	DataPosted    string
	Time          string
}

// Main script
func main() {
	requestFromFront := RequestFromFront{"", "", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/template.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if dataRequested := r.FormValue("dataRequested"); dataRequested != "" {
			requestFromFront.DataRequested = dataRequested
		}
		if dataPosted := r.FormValue("dataPosted"); dataPosted != "" {
			requestFromFront.DataPosted = dataPosted
		}

		err := scripts.RouterUrls(requestFromFront.DataRequested, requestFromFront.DataPosted, requestFromFront.Time)
		if err != nil {
			fmt.Println(err)
		}
		if err := templates.ExecuteTemplate(w, "template.html", requestFromFront); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on PORT 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
