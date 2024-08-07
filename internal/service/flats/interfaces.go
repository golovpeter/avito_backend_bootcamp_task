package flats

import "context"

type FlatsService interface {
	CreateFlat(ctx context.Context, data *CreateFlatIn) (*CreateFlatOut, error)
}
