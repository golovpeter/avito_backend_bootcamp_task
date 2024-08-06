package users

import (
	"avito_backend_bootcamp_task/internal/common"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

const registerUserQuery = `
	INSERT INTO users(email, password_hash, role)
	VALUES ($1, $2, $3)
	ON CONFLICT DO NOTHING
	RETURNING id
`

func (r *repository) CreateUser(ctx context.Context, data *UserData) (*CreateUserOut, error) {
	var newUserId int64
	row := r.conn.QueryRowContext(ctx, registerUserQuery, data.Email, data.PasswordHash, data.Role)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&newUserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, common.ErrUserAlreadyExist
	}

	return &CreateUserOut{UserID: newUserId}, nil

}

const getUserQuery = `
	SELECT id, email, password_hash, role
	FROM users
	WHERE email = $1
`

func (r *repository) GetUserData(ctx context.Context, data *UserData) (*GetUserDataOut, error) {
	var userData GetUserDataOut

	err := r.conn.GetContext(ctx, &userData, getUserQuery, data.Email)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (r *repository) GetUserRole(ctx context.Context, data *UserData) (*GetUserRoleOut, error) {
	//TODO implement me
	panic("implement me")
}
