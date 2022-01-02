package handler

import (
	Types "eduapp/CommonTypes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var SECRET_KEY = []byte("gosecretkey") //ToDo generate secret key

var generatedToken string

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GetIdByLogin(username string) (int, error) {
	db := DbConnector()
	rows, err := db.Query("SELECT id, username FROM logins WHERE username=$1", username)
	if err != nil {
		return 0, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id, &username); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
	}
	return id, nil
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func UserSignup(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var user Types.UserAuth
	var dbUser Types.UserAuth
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))

	fmt.Println(user.Password, user.Username)

	dbUser, err := GetLogin(user.Username, user.Password)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	if dbUser.Username == user.Username {
		response.Write([]byte((`{"message":"` + `trying to create existing user` + `"}"`)))
		return
	}
	CreateUserInDB(user.Username, user.Password)
	response.Write([]byte((`{"message":"` + `succesfully created user` + `"}"`)))
}

func UserLogin(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var user Types.UserAuth
	var dbUser Types.UserAuth
	json.NewDecoder(request.Body).Decode(&user)

	dbUser, err := GetLogin(user.Username, user.Password)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	generatedToken = jwtToken
	request.Header.Add("authToken", jwtToken)
	http.Redirect(response, request, "http://localhost:8080/homepage", http.StatusFound)
	return
}
