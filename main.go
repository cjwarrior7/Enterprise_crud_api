package main

import (
	"sync"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	logging "accountingService/logger"
	"accountingService/config"
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

		if err := e.Start("0.0.0.0:10000"); err != nil {
			logger.WithError(err).Fatal("Unable to start the accounting service")
		}
	})
}