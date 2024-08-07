package flats

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

const insertFlatQuery = `
	INSERT INTO flats(number, price, house_id)
	VALUES ($1, $2, $3)
	ON CONFLICT DO NOTHING 
	RETURNING id, status
`

func (r *repository) InsertNewFlat(ctx context.Context, data *InsertNewFlatIn) (*InsertNewFlatOut, error) {
	out := InsertNewFlatOut{
		HouseID: data.HouseID,
		Price:   data.Price,
		Rooms:   data.Rooms,
		Number:  data.Number,
	}

	row := r.conn.QueryRowContext(ctx, insertFlatQuery, data.Number, data.Price, data.HouseID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&out.ID, &out.Status)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, common.ErrFlatAlreadyExist
	}

	if err != nil {
		return nil, err
	}

	return &out, nil
}
