package domain

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
}

type Profile struct {
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Country   string `json:"country"`
}
