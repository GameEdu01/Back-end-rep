package main

import (
	"eduapp/pkg/db"
	"eduapp/pkg/routing"
	"encoding/json"
	"fmt"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	urlPath := ""

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	db.DbConnector()

	router := httprouter.New()
	router.GET("/about", aboutSCPR)

	fmt.Println("hello i am started")
	routing.InitRouter(router, urlPath)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}

func aboutSCPR(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	m := "Welcome to GameEduApp API"
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}
