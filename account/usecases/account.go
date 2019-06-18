package account

import (
	"context"
	"accountingService/account"
	"accountingService/models"
	"accountingService/logger"
	"fmt"
)

type accountUsecase struct {
	actRepo account.Repository
}

func (r *accountUsecase) AuthenticateUser(ctx context.Context, username string, secret string) (*models.Account, error) {
	logger.Logger.Info("Request Received into AuthenticateUser UseCase")
	fmt.Println("hello userName ", username)
	fmt.Println("hello secret ", secret)
	return r.actRepo.GetByUsername(ctx, username, secret)
}
func (r *accountUsecase) AddUser(ctx context.Context,name string,username string,user_pin string,user_email string,description string,enterprise_id string ,mobile_no string) (*models.Account, error) {
	logger.Logger.Info("Request Received into RegisterUser UseCase")
	fmt.Println("hello username ", username)
	fmt.Println("hello name ", name)
	return r.actRepo.Enterprise_AddUser(ctx,name,username,user_pin,user_email,description,enterprise_id , mobile_no)
}
func (r *accountUsecase) AddEnterprise(ctx context.Context,enterprise_email string,enterprise_username string,enterprise_pin string,enterprise_mobile string) (*models.Account, error) {
	logger.Logger.Info("Request Received into RegisterUser UseCase")
	fmt.Println("hello username ",enterprise_username)
	fmt.Println("hello email ", enterprise_email)
	return r.actRepo.AddEnterpriseUser(ctx,enterprise_email,enterprise_username,enterprise_pin,enterprise_mobile)
}
func (r *accountUsecase) GetEnterpriseUser(ctx context.Context) ([]*models.Enterprise_User, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.GetAllEnterpriseUser(ctx)
}

func (r *accountUsecase) UserofEnterprise(ctx context.Context,enterprise_id string) ([]*models.Userof_Enterprise, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.GetUserofEnterprise(ctx,enterprise_id )
}

func (r *accountUsecase) DeleteEnterprise(ctx context.Context,enterprise_id string) (*models.Account, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.DeleteEnterp(ctx,enterprise_id )
}

func (r *accountUsecase) DeleteUserofEnterprise(ctx context.Context,user_id string) (*models.Account, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.DeleteUserofEnterp(ctx,user_id )
}

func (r *accountUsecase) UploadCsv(ctx context.Context,enterprise_id string ,filedata [][]string) (*models.Account, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.UploadCsvEnterp(ctx,enterprise_id,filedata)
}
func (r *accountUsecase) Mod_enterprise(ctx context.Context,enterprise_id ,enterprise_username string ,enterprise_pin string ) (*models.Account, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.Update_enterprise(ctx,enterprise_id,enterprise_username,enterprise_pin)
}
func (r *accountUsecase) ShowOne_Enterprise(ctx context.Context,enterprise_id string) (*models.ShowJoin_Enterprise, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.ShowOne_Enterprise(ctx,enterprise_id)
}

/* func (r *accountUsecase) QrcodeEnterprise(ctx context.Context,username string , secret string) (*models.Account, error) {
	logger.Logger.Info("Request Received into GetEnterpriseUser UseCase")
	//fmt.Println("hello username ",enterprise_username)
	//fmt.Println("hello email ", enterprise_email)
	return r.actRepo.QrcodeofEnterp(ctx,username,secret)
}*/
//NewRoutesUseCase creates concrete instance of routes.Usecase
func NewAccountUseCase(repository account.Repository) account.Usecase {
	return &accountUsecase{
		actRepo: repository,
	}
}
