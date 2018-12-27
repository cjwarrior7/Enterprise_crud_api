package account


import (
	"context"
	"accountingService/models"
)

//Repository is repo layer for fetching Carrier domain objects
type Repository interface {
	GetByUsername(ctx context.Context, username string,secret string) (*models.Account, error)
}