package users

type UserData struct {
	UserID       int64  `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
}

type CreateUserOut struct {
	UserID int64 `db:"id"`
}

type GetUserDataOut struct {
	UserID       int64  `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	UserType     string `db:"role"`
}

type GetUserRoleOut struct {
	UserRole string
}
