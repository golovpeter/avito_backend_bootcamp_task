package flats

import "context"

type Repository interface {
	InsertNewFlat(ctx context.Context, data *InsertNewFlatIn) (*InsertNewFlatOut, error)
}
