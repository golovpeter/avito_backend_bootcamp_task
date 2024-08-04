package main

import (
	"avito_backend_bootcamp_task/internal/common"
	"avito_backend_bootcamp_task/internal/config"
	"avito_backend_bootcamp_task/internal/handler/login"
	"avito_backend_bootcamp_task/internal/handler/register"
	"avito_backend_bootcamp_task/internal/repository/users"
	usersservice "avito_backend_bootcamp_task/internal/service/users"
	"fmt"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		logger.Error("error to parse config file")
		return
	}

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logger.Error("error to parse logger level")
		return
	}

	logger.SetLevel(level)

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.WithError(err).Error("error to create database client")
		return
	}

	usersRepository := users.NewRepository(dbConn)

	usersService := usersservice.NewService(usersRepository, cfg.Server.JwtKey)

	registerHandler := register.NewHandler(logger, usersService)
	loginHandler := login.NewHandler(logger, usersService)

	router := gin.Default()

	router.Use(requestid.New())

	router.POST("/register", registerHandler.Register)
	router.POST("/login", loginHandler.Login)

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
