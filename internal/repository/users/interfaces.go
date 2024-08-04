package users

import "context"

type Repository interface {
	CreateUser(ctx context.Context, data *UserData) (*CreateUserOut, error)
	GetUserData(ctx context.Context, data *UserData) (*GetUserDataOut, error)
	GetUserRole(ctx context.Context, data *UserData) (*GetUserRoleOut, error)
}
