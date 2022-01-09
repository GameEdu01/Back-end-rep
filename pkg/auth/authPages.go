package auth

import (
	Types "eduapp/CommonTypes"
	myerrors "eduapp/pkg/errors"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// RegisterUserPage is responsible for sending registration page to the front-end
func RegisterUserPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/Register.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
}

//TermsAndConditions is responsible for TaC page
func TermsAndConditions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/TermsAndConditions.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
}

//TermsAndConditions is responsible for TaC page
func TermsAndConditionsForWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/TermsAndConditionsForWallet.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		myerrors.Handle500(w, r)
		return
	}
}

//LoginUserPage is responsible for sending login page to the front end
func LoginUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/Login.html")
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `" +"error":"` + err.Error() + `"}`))

			return
		}
		t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
	}
}

func ForwardToNewsFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/newsfeed", http.StatusFound)
}
