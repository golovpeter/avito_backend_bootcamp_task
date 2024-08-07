package houses

import "context"

type HousesService interface {
	CreateHouse(ctx context.Context, data *CreateHouseIn) (*CreateHouseOut, error)
}
