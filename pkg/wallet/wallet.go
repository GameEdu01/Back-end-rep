package wallet

import (
	"bytes"
	Types "eduapp/CommonTypes"
	"eduapp/pkg/auth"
	myerrors "eduapp/pkg/errors"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"html/template"
	_ "io"
	"io/ioutil"
	"net/http"
	"time"
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
	resp, err := signUpForExternal(Payload.Name, Payload.Password, Payload.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	var responseFromApi Types.ExternalResponse
	err = json.Unmarshal(resp, &responseFromApi)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	resp, err = getExternalSession(Payload.Name, Payload.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	var session string
	err = json.Unmarshal(resp, &session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	partnerJWT, err := generateJWT(session)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		myerrors.Handle500(w, r)
		return
	}

	resp, err = walletRegister(partnerJWT, session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}
	var registerResponse string
	err = json.Unmarshal(resp, &registerResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + `error parsing to json to external request` + `"}`))
		myerrors.Handle400(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(registerResponse))
}

func walletRegister(partnerJWT string, sessionID string) ([]byte, error) {
	payloadBuf := new(bytes.Buffer)
	//ToDo add url of rw
	reqSend, err := http.NewRequest("POST", "", payloadBuf)
	if err != nil {
		return nil, err
	}

	reqSend.Header.Set("X-Auth-Token", partnerJWT)
	reqSend.Header.Set("X-session-ID", sessionID)

	client := &http.Client{}
	res, err := client.Do(reqSend)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resByte, err := ioutil.ReadAll(res.Body)
	return resByte, nil
}

func getExternalSession(name string, password string) ([]byte, error) {
	password = auth.GetHash([]byte(password))
	var payload Types.ExternalRegistration
	payload.Username = name
	payload.PasswordHash = password

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(payload)
	if err != nil {
		return nil, err
	}
	reqSend, err := http.NewRequest("POST", "https://gameedu-api.herokuapp.com/api/get_session", payloadBuf)

	client := &http.Client{}
	res, err := client.Do(reqSend)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resByte, err := ioutil.ReadAll(res.Body)
	return resByte, nil
}

func signUpForExternal(name string, password string, email string) ([]byte, error) {
	password = auth.GetHash([]byte(password))
	var payload Types.ExternalRegistration
	payload.Username = name
	payload.PasswordHash = password
	payload.Email = email

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(payload)
	if err != nil {
		return nil, err
	}
	reqSend, err := http.NewRequest("POST", "https://gameedu-api.herokuapp.com/api/signup", payloadBuf)

	client := &http.Client{}
	res, err := client.Do(reqSend)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resByte, err := ioutil.ReadAll(res.Body)
	return resByte, nil
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

func generateJWT(session string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Types.Claims{
		Content: session,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(auth.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
