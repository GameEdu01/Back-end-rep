package handler

import (
	"eduapp/pkg/db"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

//MarketPage responsible for giving courses for user to sell
//Future functional
func MarketPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := template.ParseFiles("./templates/Course.html")
	if err != nil { // if there is an error
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `unable to parse request` + `"}`))
		return
	}
	id, err := strconv.Atoi(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing UUID` + `"}`))
		return
	}
	content := db.GetMarketForUser(db.DbConnector(), id)
	err = t.Execute(w, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
		return
	}
}
