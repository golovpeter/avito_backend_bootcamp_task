package register

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/common"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/service/users"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service users.UserService
}

func NewHandler(log *logrus.Logger, service users.UserService) *handler {
	return &handler{log: log, service: service}
}

func (h *handler) Register(ctx *gin.Context) {
	var in RegisterIn

	if err := ctx.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	isValid, errMsg, err := common.ValidateUserData(in.Email, in.Password)
	if err != nil {
		h.log.WithError(err).Error(errMsg)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if !isValid {
		h.log.Warn(errMsg)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	userData, err := h.service.Register(ctx, &users.UserDataIn{
		Email:    in.Email,
		Password: in.Password,
		UserRole: in.UserType,
	})
	if err != nil {
		h.log.WithError(err).Error("error register user")
		ctx.JSON(http.StatusInternalServerError, common.Err5xx{
			Message:   "error register user",
			RequestID: requestid.Get(ctx),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, RegisterOut{
		UserID: userData.UserID,
	})
}
