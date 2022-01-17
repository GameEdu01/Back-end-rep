package wallet

import (
	"bytes"
	Types "eduapp/CommonTypes"
	myerrors "eduapp/pkg/errors"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	req := &Types.RequestWalletSignUpResived{}
	err = json.Unmarshal(body, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
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
		myerrors.Handle400(w, r)
		return
	}

	//Todo request with register to /api/signup
	//ToDo create session /api/new_session
	//ToDo get session /api/get_session
	//ToDo send request for wallet registration
	reqSend, err := http.NewRequest("POST", "https://gameedu-api.herokuapp.com/api/wallet_signup", payloadBuf)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	client := &http.Client{}
	res, err := client.Do(reqSend)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + `error sending request` + `"}`))
		myerrors.Handle500(w, r)
	}

	defer res.Body.Close()

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing response` + `"}`))
		myerrors.Handle400(w, r)
	}

	msg, err := (io.ReadAll(res.Body))
	fmt.Println(string(msg))
	w.WriteHeader(http.StatusOK)
	w.Write(resByte)
}

func WalletVerifyPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/CreateWallet.html")
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `" +"error":"` + err.Error() + `"}`))
			myerrors.Handle500(w, r)
			return
		}
		err = t.Execute(w, nil)
		if err != nil { // if there is an error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + `template parsing error` + `"}`))
			myerrors.Handle500(w, r)
			return
		}
	}
}
