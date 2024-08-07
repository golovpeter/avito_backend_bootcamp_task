package houses

import (
	"avito_backend_bootcamp_task/internal/repository/houses"
	"context"
)

type service struct {
	repository houses.Repository
}

func NewService(repository houses.Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateHouse(ctx context.Context, data *CreateHouseIn) (*CreateHouseOut, error) {
	houseData, err := s.repository.InsertNewHouse(ctx, &houses.InsertNewHouseIn{
		Address:   data.Address,
		Year:      data.Year,
		Developer: data.Developer,
	})
	if err != nil {
		return nil, err
	}

	return &CreateHouseOut{
		ID:        houseData.ID,
		Address:   houseData.Address,
		Year:      houseData.Year,
		Developer: houseData.Developer,
		CreatedAt: houseData.CreatedAt,
		UpdatedAt: houseData.UpdatedAt,
	}, nil
}
