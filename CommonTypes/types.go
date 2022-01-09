package CommonTypes

import (
	"github.com/dgrijalva/jwt-go"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Content struct {
	Request string `json:"paragraph_text"`
}

type Course struct {
	Id          int    `json:"id"`
	Author_id   int    `json:"author_id"`
	Category    string `json:"category"`
	Game        string `json:"game"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Views       int    `json:"views"`
}

type ResivedCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Game        string `json:"game"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Agree       string `json:"agree"`
}

type RequestWalletSignUpSend struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
	Password string `json:"password"`
}

type RequestWalletSignUpResived struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Agree    string `json:"agree"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PageVariables struct {
	Date string
	Time string
}

// Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
