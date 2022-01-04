package CommonTypes

import "github.com/dgrijalva/jwt-go"

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
	Request string `json:"request"`
}

type Course struct {
	Id             int     `json:"id"`
	Author_id      int     `json:"author_id"`
	Price          float64 `json:"price"`
	Game_name      string  `json:"game_name"`
	Followers      int     `json:"followers"`
	Course_content Content `json:"course_content"`
	Owners         []int   `json:"owners"`
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

type RequestCourse struct {
	Price          string  `json:"price"`
	Game_name      string  `json:"game_name"`
	Course_content Content `json:"course_content"`
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
