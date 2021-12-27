package scripts

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var SECRET_KEY = []byte("gosecretkey") //ToDo generate secret key

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var generatedToken string

var client *mongo.Client

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

func UserSignup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user UserAuth
	var dbUser UserAuth
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

func UserLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user UserAuth
	var dbUser UserAuth
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
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
	generatedToken = jwtToken
	return
}

func VerifyTokens(authToken string) bool {
	if authToken == generatedToken {
		return true
	}
	return false
}
