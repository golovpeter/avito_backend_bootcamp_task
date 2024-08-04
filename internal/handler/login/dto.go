package login

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOut struct {
	Token string `json:"token"`
}
