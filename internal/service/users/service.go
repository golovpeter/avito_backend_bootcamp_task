package users

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/repository/users"
	"context"
)

type service struct {
	repository users.Repository

	jwtKey string
}

func NewService(repository users.Repository, jwtKey string) *service {
	return &service{repository: repository, jwtKey: jwtKey}
}

func (s *service) Register(ctx context.Context, data *UserDataIn) (*RegisterOut, error) {
	passwordHash := common.GeneratePasswordHash(data.Password)

	repoData, err := s.repository.CreateUser(ctx, &users.UserData{
		Email:        data.Email,
		PasswordHash: passwordHash,
		Role:         data.UserRole,
	})
	if err != nil {
		return nil, err
	}

	return &RegisterOut{UserID: repoData.UserID}, nil
}

func (s *service) Login(ctx context.Context, data *UserDataIn) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) Authorization(ctx context.Context, data *UserDataIn) (*AuthorizationOut, error) {
	//TODO implement me
	panic("implement me")
}
