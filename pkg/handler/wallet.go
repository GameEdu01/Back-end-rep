package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func WalletVerifyPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/createWallet.html")
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
		} else {
			username := r.Form["username"]
			password := r.Form["password"]
			fmt.Fprintf(w, "username = %s, password = %s", username, password)
		}
	}
}
