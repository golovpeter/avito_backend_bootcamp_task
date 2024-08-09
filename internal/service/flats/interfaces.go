package flats

import "context"

type FlatsService interface {
	CreateFlat(ctx context.Context, data *CreateFlatIn) (*FlatData, error)
	UpdateFlatStatus(ctx context.Context, data *UpdateFlatIn) (*FlatData, error)
	GetFlatsByHouseID(ctx context.Context, data *GetFlatsByHouseID) ([]*FlatData, error)
}
