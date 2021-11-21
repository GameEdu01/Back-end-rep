package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type req struct {
	Data        string `json:"data"`
	Description string `json:"description"`
}

var requests []req

// getHandler
func getHandler(w http.ResponseWriter, r *http.Request) {
	reqListBytes, err := json.Marshal(requests)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(reqListBytes)
}

// createHandler
func createHandler(w http.ResponseWriter, r *http.Request) {
	request := req{
		Data:        "data",
		Description: "description",
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request.Data = r.Form.Get("data")
	request.Description = r.Form.Get("description")

	requests = append(requests, request)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
