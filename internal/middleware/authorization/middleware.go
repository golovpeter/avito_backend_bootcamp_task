package authorization

import (
	"avito_backend_bootcamp_task/internal/common"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Authorization(
	logger *logrus.Logger,
	enforcer *casbin.Enforcer,
	jwtKey string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")

		if accessToken == "" {
			logger.Error(common.ErrInvalidAuthHeader)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(accessToken, " ")

		if splitToken[0] != "Bearer" {
			logger.Error(common.ErrInvalidAuthHeader)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := common.GetTokenClaims(splitToken[1], jwtKey)
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		allow, err := enforcer.Enforce(
			claims["UserType"].(string),
			ctx.Request.URL.Path,
			ctx.Request.Method,
		)
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.Err5xx{
				Message:   "failed authorize user",
				RequestID: requestid.Get(ctx),
				Code:      http.StatusInternalServerError,
			})
			return
		}

		if !allow {
			logger.Error(common.ErrAccessDenied)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Set("user_type", claims["UserType"].(string))
		ctx.Next()
	}
}
