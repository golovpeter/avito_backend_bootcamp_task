package users

import "context"

type UserService interface {
	Register(ctx context.Context, data *UserDataIn) (*RegisterOut, error)
	Login(ctx context.Context, data *UserDataIn) error
	Authorization(ctx context.Context, data *UserDataIn) (*AuthorizationOut, error)
}
