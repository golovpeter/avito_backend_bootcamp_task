package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/common"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/config"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/flat_create"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/get_flats"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/house_create"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/login"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/register"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/handler/update_flat_status"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/middleware/authorization"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/repository/flats"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/repository/houses"
	"github.com/golovpeter/avito_backend_bootcamp_task/internal/repository/users"
	flatsservice "github.com/golovpeter/avito_backend_bootcamp_task/internal/service/flats"
	housesservice "github.com/golovpeter/avito_backend_bootcamp_task/internal/service/houses"
	usersservice "github.com/golovpeter/avito_backend_bootcamp_task/internal/service/users"
	"github.com/sirupsen/logrus"
)

const (
	casbinModelPath  = "internal/config/casbin_config/model.conf"
	casbinPolicyPath = "internal/config/casbin_config/policy.csv"
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
	housesRepository := houses.NewRepository(dbConn)
	flatsRepository := flats.NewRepository(dbConn)

	usersService := usersservice.NewService(usersRepository, cfg.Server.JwtKey)
	housesService := housesservice.NewService(housesRepository)
	flatsService := flatsservice.NewService(flatsRepository)

	registerHandler := register.NewHandler(logger, usersService)
	loginHandler := login.NewHandler(logger, usersService)
	createHouseHandler := house_create.NewHandler(logger, housesService)
	createFlatHandler := flat_create.NewHandler(logger, flatsService)
	updateFlatStatusHandler := update_flat_status.NewHandler(logger, flatsService)
	getFlatsHandler := get_flats.NewHandler(logger, flatsService)

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
		houseGroup.POST("/create", createHouseHandler.CreateHouse)
		houseGroup.GET("/:id", getFlatsHandler.GetFlats)
	}

	flatGroup := router.Group("/flat").Use(
		authorization.Authorization(logger, enforcer, cfg.Server.JwtKey))
	{
		flatGroup.POST("/create", createFlatHandler.CreateFlat)
		flatGroup.POST("/update", updateFlatStatusHandler.UpdateFlatStatus)
	}

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
