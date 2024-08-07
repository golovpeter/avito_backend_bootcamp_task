package houses

import "context"

type Repository interface {
	InsertNewHouse(ctx context.Context, data *InsertNewHouseIn) (*InsertNewHouseOut, error)
}
