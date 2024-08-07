package houses

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

func NewRepository(dbConn *sqlx.DB) *repository {
	return &repository{conn: dbConn}
}

const insertHouseQuery = `
	INSERT INTO houses(address, year, developer)
	VALUES ($1, $2, $3)
	ON CONFLICT DO NOTHING 
	RETURNING id, created_at, updated_at
`

func (r *repository) InsertNewHouse(ctx context.Context, data *InsertNewHouseIn) (*InsertNewHouseOut, error) {
	out := InsertNewHouseOut{
		Address:   data.Address,
		Year:      data.Year,
		Developer: data.Developer,
	}

	row := r.conn.QueryRowContext(ctx, insertHouseQuery, data.Address, data.Year, data.Developer)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&out.ID, &out.CreatedAt, &out.UpdatedAt)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, common.ErrHouseAlreadyExist
	}

	if err != nil {
		return nil, err
	}

	return &out, nil
}
