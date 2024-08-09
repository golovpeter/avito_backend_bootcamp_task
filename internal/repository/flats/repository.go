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
	INSERT INTO flats(number, price, house_id, rooms)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT DO NOTHING 
	RETURNING id, status
`

func (r *repository) InsertNewFlat(ctx context.Context, data *InsertNewFlatIn) (*FlatData, error) {
	out := FlatData{
		HouseID: data.HouseID,
		Price:   data.Price,
		Rooms:   data.Rooms,
		Number:  data.Number,
	}

	row := r.conn.QueryRowContext(ctx, insertFlatQuery, data.Number, data.Price, data.HouseID, data.Rooms)
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

const updateFlatStatusQuery = `
	UPDATE flats 
	SET status = $1
	WHERE id = $2
	RETURNING house_id, price, rooms, number
`

func (r *repository) UpdateFlatStatus(ctx context.Context, data *UpdateFlatIn) (*FlatData, error) {
	out := FlatData{
		ID:     data.ID,
		Status: data.Status,
	}

	row := r.conn.QueryRowContext(ctx, updateFlatStatusQuery, data.Status, data.ID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&out.HouseID, &out.Price, &out.Rooms, &out.Number)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, common.ErrUpdateFlatStatus
	}

	if err != nil {
		return nil, err
	}

	return &out, nil
}

const getFlatsClientQuery = `
	SELECT *
	FROM flats
	WHERE house_id = $1
`

func (r *repository) GetFlatsByHouseID(ctx context.Context, data *GetFlatsIn) ([]*FlatData, error) {
	out := make([]*FlatData, 0)

	rows, err := r.conn.QueryxContext(ctx, getFlatsClientQuery, data.HouseID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var flat FlatData
		if err = rows.Scan(
			&flat.ID,
			&flat.Number,
			&flat.Rooms,
			&flat.Price,
			&flat.HouseID,
			&flat.Status,
		); err != nil {
			return nil, err
		}

		out = append(out, &flat)
	}

	return out, nil
}
