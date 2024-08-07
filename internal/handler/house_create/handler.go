package house_create

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/service/houses"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service houses.HousesService
}

func NewHandler(logger *logrus.Logger, service houses.HousesService) *handler {
	return &handler{log: logger, service: service}
}

func (h *handler) CreateHouse(ctx *gin.Context) {
	var in CreateHouseIn

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

	houseData, err := h.service.CreateHouse(ctx, &houses.CreateHouseIn{
		Address:   in.Address,
		Year:      in.Year,
		Developer: in.Developer,
	})
	if err != nil {
		h.log.WithError(err).Error("error create house")
		ctx.JSON(http.StatusInternalServerError, common.Err5xx{
			Message:   "error create house",
			RequestID: requestid.Get(ctx),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	out := &CreateHouseOut{
		ID:        houseData.ID,
		Address:   houseData.Address,
		Year:      houseData.Year,
		Developer: houseData.Developer,
		CreatedAt: houseData.CreatedAt,
		UpdatedAt: houseData.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, out)
}
