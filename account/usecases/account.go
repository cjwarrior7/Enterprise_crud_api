package account

import (
	"context"
	"accountingService/account"
	"accountingService/models"
	"accountingService/logger"
)

type accountUsecase struct {
	actRepo account.Repository
}

func (r *accountUsecase) AuthenticateUser(ctx context.Context, username string, secret string) (*models.Account, error) {
	logger.Logger.Info("Request Received into AuthenticateUser UseCase")
	return r.actRepo.GetByUsername(ctx, username, secret)
}

//NewRoutesUseCase creates concrete instance of routes.Usecase
func NewAccountUseCase(repository account.Repository) account.Usecase {
	return &accountUsecase{
		actRepo: repository,
	}
}
