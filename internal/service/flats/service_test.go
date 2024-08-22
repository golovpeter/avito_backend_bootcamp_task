package flats

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/repository/flats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockFlatsRepository *flats.MockRepository

	service FlatsService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockFlatsRepository = flats.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockFlatsRepository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var (
	testID      int64 = 1
	testHouseID int64 = 1
	testPrice   int64 = 10000
	testRooms         = 1
	testNumber        = 10
)

func (ts *TestSuite) Test_CreateFlat_Success() {
	successResponse := &flats.FlatData{
		HouseID: testHouseID,
		Price:   testPrice,
		Rooms:   testRooms,
		Number:  testNumber,
		Status:  "created",
	}

	ts.mockFlatsRepository.EXPECT().
		InsertNewFlat(context.TODO(), &flats.InsertNewFlatIn{
			HouseID: testHouseID,
			Price:   testPrice,
			Rooms:   testRooms,
			Number:  testNumber,
		}).
		Times(1).
		Return(successResponse, nil)

	flatData, err := ts.service.CreateFlat(context.TODO(), &CreateFlatIn{
		HouseID: testHouseID,
		Price:   testPrice,
		Rooms:   testRooms,
		Number:  testNumber,
	})

	assert.NoError(ts.T(), err)
	assert.EqualValues(ts.T(), successResponse, flatData)
}

func (ts *TestSuite) Test_CreateFlat_RepositoryError() {
	ts.mockFlatsRepository.EXPECT().
		InsertNewFlat(context.TODO(), &flats.InsertNewFlatIn{
			HouseID: testHouseID,
			Price:   testPrice,
			Rooms:   testRooms,
			Number:  testNumber,
		}).
		Times(1).
		Return(nil, errors.New("insert error"))

	flatData, err := ts.service.CreateFlat(context.TODO(), &CreateFlatIn{
		HouseID: testHouseID,
		Price:   testPrice,
		Rooms:   testRooms,
		Number:  testNumber,
	})

	var nilFLat *flats.FlatData
	assert.Error(ts.T(), err)
	assert.EqualValues(ts.T(), nilFLat, flatData)
}

func (ts *TestSuite) Test_GetFlatsByHouseID_SuccessClient() {
	successResponse := []*flats.FlatData{
		{
			ID:      testID,
			HouseID: testHouseID,
			Price:   testPrice,
			Rooms:   testRooms,
			Status:  "approved",
			Number:  testNumber,
		},
	}

	ts.mockFlatsRepository.EXPECT().
		GetFlatsByHouseID(context.TODO(), &flats.GetFlatsIn{
			HouseID: testHouseID,
		}).
		Times(1).
		Return(successResponse, nil)

	flatsInfo, err := ts.service.GetFlatsByHouseID(context.TODO(), &GetFlatsByHouseID{
		HouseID:  testHouseID,
		UserType: "client",
	})

	assert.NoError(ts.T(), err)

	for i := 0; i < len(successResponse); i++ {
		assert.EqualValues(ts.T(), successResponse[i], flatsInfo[i])
	}
}

func (ts *TestSuite) Test_GetFlatsByHouseID_SuccessModerator() {
	successResponse := []*flats.FlatData{
		{
			ID:      testID,
			HouseID: testHouseID,
			Price:   testPrice,
			Rooms:   testRooms,
			Status:  "created",
			Number:  testNumber,
		},
	}

	ts.mockFlatsRepository.EXPECT().
		GetFlatsByHouseID(context.TODO(), &flats.GetFlatsIn{
			HouseID: testHouseID,
		}).
		Times(1).
		Return(successResponse, nil)

	flatsInfo, err := ts.service.GetFlatsByHouseID(context.TODO(), &GetFlatsByHouseID{
		HouseID:  testHouseID,
		UserType: "moderator",
	})

	assert.NoError(ts.T(), err)

	for i := 0; i < len(successResponse); i++ {
		assert.EqualValues(ts.T(), successResponse[i], flatsInfo[i])
	}
}

func (ts *TestSuite) Test_GetFlatsByHouseID_RepositoryError() {
	ts.mockFlatsRepository.EXPECT().
		GetFlatsByHouseID(context.TODO(), &flats.GetFlatsIn{
			HouseID: testHouseID,
		}).
		Times(1).
		Return(nil, errors.New("repository error"))

	flatsInfo, err := ts.service.GetFlatsByHouseID(context.TODO(), &GetFlatsByHouseID{
		HouseID:  testHouseID,
		UserType: "moderator",
	})

	var nilFLat []*FlatData
	assert.Error(ts.T(), err)
	assert.Equal(ts.T(), flatsInfo, nilFLat)
}
