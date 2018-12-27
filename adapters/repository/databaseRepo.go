package repository

import (
	"database/sql"
	"fmt"

	//For initialization
	_ "github.com/lib/pq"
	config "accountingService/config"
	"accountingService/logger"
)

// NewDBAdapterRepository - Repository layer for database connection
func NewDBAdapterRepository(config *config.Config) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Pass, config.Database.DBName)

	dbConn, err := sql.Open(config.Database.DBType, connectionString)
	if err != nil {
		logger.Logger.WithError(err).WithField("connection_string", connectionString).Errorf("Unable to connect to database")
		return nil
	}
	logger.Logger.WithField("connection_string", connectionString).Info("connect to database")
	dbConn.SetMaxOpenConns(config.Database.PoolSize)
	dbConn.SetMaxIdleConns(config.Database.PoolSize)
	return dbConn
}
