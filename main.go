package main

import (
	"sync"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	logging "accountingService/logger"
	"accountingService/config"
	adapterRepository "accountingService/adapters/repository"
	accountingRepo "accountingService/account/repository"
	accountingUsecase "accountingService/account/usecases"
	accountingController "accountingService/account/controller"
)

var onceRest sync.Once

func main() {
	onceRest.Do(func() {
		e := echo.New()
		//Setting up the config
		config := configs.GetConfig()
		//Setting up the Logger
		logger := logging.NewLogger(config.Log.LogFile, config.Log.LogFile)

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORS())
		
		dbAdapter := adapterRepository.NewDBAdapterRepository(config)
		accountRepo := accountingRepo.NewAccountRepository(dbAdapter)
		accountUc := accountingUsecase.NewAccountUseCase(accountRepo)
		accountingController.NewAccountController(e,accountUc)

		if err := e.Start("0.0.0.0:10000"); err != nil {
			logger.WithError(err).Fatal("Unable to start the accounting service")
		}
	})
}