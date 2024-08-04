package register

type RegisterIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type RegisterOut struct {
	UserID int64 `json:"user_id"`
}
