package handler

import (
	"bytes"
	Types "eduapp/CommonTypes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"net/http"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}

	req := &Types.RequestWalletSignUpResived{}
	err = json.Unmarshal(body, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}
	fmt.Println(*req)

	var Payload Types.RequestWalletSignUpSend
	Payload.Name = req.Name
	Payload.Surname = req.Surname
	Payload.Email = req.Email
	Payload.Phone = req.Phone
	Payload.Password = req.Password

	payloadBuf := new(bytes.Buffer)
	err = json.NewEncoder(payloadBuf).Encode(Payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request to json` + `"}`))
		return
	}

	reqSend, err := http.NewRequest("POST", "https://gameedu-api.herokuapp.com/api/wallet_signup", payloadBuf)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		return
	}

	client := &http.Client{}
	res, err := client.Do(reqSend)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `error sending request` + `"}`))
	}

	defer res.Body.Close()

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing response` + `"}`))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resByte)
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
		err = t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			return
		}
	}
}
