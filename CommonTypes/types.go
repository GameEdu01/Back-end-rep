package CommonTypes

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	id       string
	username string
	password string
}

type Content struct {
	Request string `json:"request"`
}

type Course struct {
	Id             int     `json:"id"`
	Author_id      int     `json:"author_id"`
	Price          float64 `json:"price"`
	Owners         []int   `json:"owners"`
	Game_name      string  `json:"game_name"`
	Followers      int     `json:"followers"`
	Course_content Content `json:"course_content"`
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
