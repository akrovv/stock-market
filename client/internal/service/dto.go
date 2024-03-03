package service

// user
type SaveUser struct {
	Email    string
	Nickname string
	Password string
	Country  string
}

type GetUser struct {
	Email    string
	Password string
}

// session
type GetSession struct {
	ID string
}
