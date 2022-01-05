package auth

import (
	Types "eduapp/CommonTypes"
	db2 "eduapp/pkg/db"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GetIdByLogin(username string) (int, error) {
	db := db2.DbConnector()
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

func UserSignup(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var user Types.UserAuth
	var dbUser Types.User
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))

	fmt.Println(user.Password, user.Username)

	dbUser, err := db2.GetLogin(user.Username)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	if dbUser.Username == user.Username {
		response.Write([]byte((`{"message":"` + `trying to create existing user` + `"}"`)))
		return
	}
	db2.CreateUserInDB(user.Username, user.Password)
	response.Write([]byte((`{"message":"` + `succesfully created user` + `"}"`)))
}

func UserLogin(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var user Types.UserAuth
	var dbUser Types.User
	json.NewDecoder(request.Body).Decode(&user)

	dbUser, err := db2.GetLogin(user.Username)

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

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	var creds Types.Credentials
	claims := &Types.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"authToken":"` + tokenString + `"}`))
	fmt.Println(tokenString)
	return
}
