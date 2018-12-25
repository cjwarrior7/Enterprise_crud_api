package account

import (
"context"
"accountingService/account"
"accountingService/models"
)

type accountUsecase struct {
	actRepo account.Repository
}

func (r *accountUsecase) AuthenticateUser(ctx context.Context, username string) (*models.Account, error) {
	return r.actRepo.GetByUsername(ctx,username)
}

//NewRoutesUseCase creates concrete instance of routes.Usecase
func NewAccountUseCase(repository account.Repository) account.Usecase {
	return &accountUsecase{
		actRepo: repository,
	}
}