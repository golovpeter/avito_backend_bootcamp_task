package flats

import (
	"avito_backend_bootcamp_task/internal/repository/flats"
	"context"
)

type service struct {
	repository flats.Repository
}

func NewService(repository flats.Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateFlat(ctx context.Context, data *CreateFlatIn) (*FlatData, error) {
	flatData, err := s.repository.InsertNewFlat(ctx, &flats.InsertNewFlatIn{
		HouseID: data.HouseID,
		Price:   data.Price,
		Rooms:   data.Rooms,
		Number:  data.Number,
	})
	if err != nil {
		return nil, err
	}

	return &FlatData{
		ID:      flatData.ID,
		HouseID: flatData.HouseID,
		Price:   flatData.Price,
		Rooms:   flatData.Rooms,
		Number:  flatData.Number,
		Status:  flatData.Status,
	}, nil
}

func (s *service) UpdateFlatStatus(ctx context.Context, data *UpdateFlatIn) (*FlatData, error) {
	flatData, err := s.repository.UpdateFlatStatus(ctx, &flats.UpdateFlatIn{
		ID:     data.ID,
		Status: data.Status,
	})
	if err != nil {
		return nil, err
	}

	return &FlatData{
		ID:      flatData.ID,
		HouseID: flatData.HouseID,
		Price:   flatData.Price,
		Rooms:   flatData.Rooms,
		Number:  flatData.Number,
		Status:  flatData.Status,
	}, nil

}
