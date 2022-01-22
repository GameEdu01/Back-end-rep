package auth

import (
	Types "eduapp/CommonTypes"
	myerrors "eduapp/pkg/errors"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// SignAndLogin is responsible for sending registration page to the front-end
func SignAndLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	HomePageVars := Types.PageVariables{}
	t, err := template.ParseFiles("./templates/auth.html")
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

// TermsAndConditionsForWallet is responsible for TaC page
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

func ForwardToNewsFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/newsfeed", http.StatusFound)
}
