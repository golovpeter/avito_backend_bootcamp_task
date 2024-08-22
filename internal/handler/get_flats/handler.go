package get_flats

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/common"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/service/flats"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service flats.FlatsService
}

func NewHandler(log *logrus.Logger, service flats.FlatsService) *handler {
	return &handler{log: log, service: service}
}

func (h *handler) GetFlats(ctx *gin.Context) {
	houseIDParam := ctx.Param("id")

	houseID, err := strconv.Atoi(houseIDParam)
	if err != nil {
		h.log.WithError(err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, errMsg, err := validateInParams(houseID)
	if err != nil {
		h.log.WithError(err).Error(err.Error())
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isValid {
		h.log.Error(errMsg)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userType, exist := ctx.Get("user_type")
	if !exist {
		h.log.Error(common.ErrUserTypeNotExist)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	serviceFlats, err := h.service.GetFlatsByHouseID(ctx, &flats.GetFlatsByHouseID{
		HouseID:  int64(houseID),
		UserType: userType.(string),
	})
	if err != nil {
		h.log.WithError(err).Error("error get flats")
		ctx.JSON(http.StatusInternalServerError, common.Err5xx{
			Message:   "error get serviceFlats",
			RequestID: requestid.Get(ctx),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	outFlats := make([]*FlatData, 0)
	for _, serviceFlat := range serviceFlats {
		outFlats = append(outFlats, &FlatData{
			ID:      serviceFlat.ID,
			HouseID: serviceFlat.HouseID,
			Price:   serviceFlat.Price,
			Rooms:   serviceFlat.Rooms,
			Number:  serviceFlat.Number,
			Status:  serviceFlat.Status,
		})
	}

	out := GetFlatsOut{
		Flats: outFlats,
	}

	ctx.JSON(http.StatusOK, out)
}
