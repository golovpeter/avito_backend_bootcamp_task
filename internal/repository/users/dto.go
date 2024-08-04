package users

type UserData struct {
	Email        string
	PasswordHash string
	Role         string
}

type CreateUserOut struct {
	UserID int64
}

type GetUserDataOut struct {
	PasswordHash string
}

type GetUserRoleOut struct {
	UserRole string
}
