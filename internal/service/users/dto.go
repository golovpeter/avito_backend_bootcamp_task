package users

type UserDataIn struct {
	Email    string
	Password string
	UserRole string
}

type RegisterOut struct {
	UserID int64
}

type AuthorizationOut struct {
	Token string
}
