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

func (s *service) GetFlatsByHouseID(ctx context.Context, data *GetFlatsByHouseID) ([]*FlatData, error) {
	var out []*FlatData

	flats, err := s.repository.GetFlatsByHouseID(ctx, &flats.GetFlatsIn{
		HouseID: data.HouseID,
	})
	if err != nil {
		return nil, err
	}

	switch data.UserType {
	case "moderator":
		for _, flat := range flats {
			out = append(out, &FlatData{
				ID:      flat.ID,
				HouseID: flat.HouseID,
				Price:   flat.Price,
				Rooms:   flat.Rooms,
				Number:  flat.Number,
				Status:  flat.Status,
			})
		}
	case "client":
		for _, flat := range flats {
			if flat.Status == "approved" {
				out = append(out, &FlatData{
					ID:      flat.ID,
					HouseID: flat.HouseID,
					Price:   flat.Price,
					Rooms:   flat.Rooms,
					Number:  flat.Number,
					Status:  flat.Status,
				})
			}
		}
	}

	return out, nil
}
