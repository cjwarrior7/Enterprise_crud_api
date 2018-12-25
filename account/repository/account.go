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
	Conn *sql.DB
}

// NewCarrierRepository - Repository for carrier
func NewAccountRepository(Conn *sql.DB)  account.Repository{
	return &accountRepository{
		Conn: Conn,
	}
}

func  (c *accountRepository) GetByUsername(ctx context.Context, username string) (*models.Account, error) {
	query := fmt.Sprintf("SELECT * FROM carriers WHERE UserName = %s", username)
	rows, err := c.Conn.QueryContext(ctx, query)
	if err != nil {
		logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	result := make([]*models.Account, 0)
	for rows.Next() {
		t := new(models.Account)
		err := rows.Scan(&t.ID, &t.Name)
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
