package account

import (
"context"
"accountingService/models"
)

//Usecase for routes
type Usecase interface {
	AuthenticateUser(ctx context.Context, username string) (*models.Account, error)
}