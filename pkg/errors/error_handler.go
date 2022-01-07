package errors

import (
	Types "eduapp/CommonTypes"
	"html/template"
	"net/http"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/errors/error404.html")
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

func Handle401(w http.ResponseWriter, r *http.Request) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/errors/error401.html")
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

func Handle400(w http.ResponseWriter, r *http.Request) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/errors/error400.html")
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

func Handle500(w http.ResponseWriter, r *http.Request) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/errors/error500.html")
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
