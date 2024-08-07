package flat_create

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/service/flats"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service flats.FlatsService
}

func NewHandler(log *logrus.Logger, service flats.FlatsService) *handler {
	return &handler{log: log, service: service}
}

func (h *handler) CreateFlat(ctx *gin.Context) {
	var in CreateFlatIn

	if err := ctx.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, errMsg, err := validateInParams(in)
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

	flatData, err := h.service.CreateFlat(ctx, &flats.CreateFlatIn{
		HouseID: in.HouseID,
		Price:   in.Price,
		Rooms:   in.Rooms,
		Number:  in.Number,
	})
	if err != nil {
		h.log.WithError(err).Error("error create flat")
		ctx.JSON(http.StatusInternalServerError, common.Err5xx{
			Message:   "error create flat",
			RequestID: requestid.Get(ctx),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	out := &CreateFlatOut{
		ID:      flatData.ID,
		HouseID: flatData.HouseID,
		Price:   flatData.Price,
		Rooms:   flatData.Rooms,
		Status:  flatData.Status,
		Number:  flatData.Number,
	}

	ctx.JSON(http.StatusOK, out)
}
