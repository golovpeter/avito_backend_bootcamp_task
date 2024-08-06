package main

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/config"
	"avito_backend_bootcamp_task/internal/handler/login"
	"avito_backend_bootcamp_task/internal/handler/register"
	"avito_backend_bootcamp_task/internal/middleware/authorization"
	"avito_backend_bootcamp_task/internal/repository/users"
	usersservice "avito_backend_bootcamp_task/internal/service/users"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	casbinModelPath  = "casbin_configs/model.conf"
	casbinPolicyPath = "casbin_configs/policy.csv"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		logger.Error("error to parse config file: " + err.Error())
		return
	}

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logger.Error("error to parse logger level: " + err.Error())
		return
	}

	logger.SetLevel(level)

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.WithError(err).Error("error to create database client: " + err.Error())
		return
	}

	enforcer, err := casbin.NewEnforcer(casbinModelPath, casbinPolicyPath)
	if err != nil {
		logger.Error("error create enforcer: " + err.Error())
		return
	}

	usersRepository := users.NewRepository(dbConn)

	usersService := usersservice.NewService(usersRepository, cfg.Server.JwtKey)

	registerHandler := register.NewHandler(logger, usersService)
	loginHandler := login.NewHandler(logger, usersService)

	router := gin.Default()
	router.Use(requestid.New())

	public := router.Group("")
	{
		public.POST("/login", loginHandler.Login)
		public.POST("/register", registerHandler.Register)
	}

	houseGroup := router.Group("/house").Use(
		authorization.Authorization(logger, enforcer, cfg.Server.JwtKey))
	{
		houseGroup.POST("/create", func(context *gin.Context) {})
	}

	flatGroup := router.Group("/flat").Use(
		authorization.Authorization(logger, enforcer, cfg.Server.JwtKey))
	{
		flatGroup.POST("/create", func(context *gin.Context) {})
	}

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
