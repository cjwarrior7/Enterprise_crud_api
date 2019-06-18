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
		fmt.Printf("hi:%s",connectionString)

	dbConn, err := sql.Open("postgres", connectionString)//config.Database.DBType
	if err != nil {
		fmt.Println("hi erorr:",err)
		logger.Logger.WithError(err).WithField("connection_string", connectionString).Errorf("Unable to connect to database")
		return nil
	}
	logger.Logger.WithField("connection_string", connectionString).Info("connect to database")
	dbConn.SetMaxOpenConns(config.Database.PoolSize)
	dbConn.SetMaxIdleConns(config.Database.PoolSize)
	return dbConn
}
