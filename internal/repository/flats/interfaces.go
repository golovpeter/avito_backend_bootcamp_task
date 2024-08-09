package flats

import "context"

type Repository interface {
	InsertNewFlat(ctx context.Context, data *InsertNewFlatIn) (*FlatData, error)
	UpdateFlatStatus(ctx context.Context, data *UpdateFlatIn) (*FlatData, error)
	GetFlatsByHouseID(ctx context.Context, data *GetFlatsIn) ([]*FlatData, error)
}
