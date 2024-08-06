package login

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/service/users"
	"errors"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service users.UserService
}

func NewHandler(log *logrus.Logger, service users.UserService) *handler {
	return &handler{log: log, service: service}
}

func (h *handler) Login(ctx *gin.Context) {
	var in LoginIn

	if err := ctx.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, errMsg, err := common.ValidateUserData(in.Email, in.Password)
	if err != nil {
		h.log.WithError(err).Error(errMsg)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isValid {
		h.log.Warn(errMsg)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(ctx, &users.UserDataIn{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil && errors.Is(common.ErrUserNotExist, err) {
		h.log.WithError(err).Error("user does not exist")
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		h.log.WithError(err).Error("error login user")
		ctx.JSON(http.StatusInternalServerError, common.Err5xx{
			Message:   "error login user",
			RequestID: requestid.Get(ctx),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, &LoginOut{
		Token: token,
	})
}
