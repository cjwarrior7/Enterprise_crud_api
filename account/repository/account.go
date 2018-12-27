package repository

import (
	"context"
	"database/sql"
	"fmt"
	"accountingService/logger"
	"accountingService/models"
	"accountingService/account"
)

type accountRepository struct {
	DbConn *sql.DB
}

// NewCarrierRepository - Repository for carrier
func NewAccountRepository(Conn *sql.DB)  account.Repository{
	return &accountRepository{
		DbConn: Conn,
	}
}

func  (c *accountRepository) GetByUsername(ctx context.Context, username string) (*models.Account, error) {
	query := fmt.Sprintf("SELECT id,is_superuser FROM accounts_login WHERE is_active=true and username = '%s'", username)
	rows, err := c.DbConn.QueryContext(ctx, query)
	if err != nil {
		logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching account_details from database")
		return nil, err
	}
	logger.Logger.Info("In accountRepo:GetByUsername")
	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	result := make([]*models.Account, 0)
	for rows.Next() {
		t := new(models.Account)
		err := rows.Scan(&t.ID, &t.SuperUser)
		if err != nil {
			logger.Logger.WithError(err).WithField("query", query).
				Errorf("Error while fetching carriers from database")
			return nil, err
		}
		result = append(result, t)
	}

	if len(result) == 0 {
		logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
		return nil, fmt.Errorf("carrier not found")
	}
	return result[0], nil
}
