package account

import (
"context"
"accountingService/models"
)

//Usecase for routes
type Usecase interface {
	AuthenticateUser(ctx context.Context, username string, secret string) (*models.Account, error)
	AddUser(ctx context.Context,name string, username string,user_pin string,user_email string,description string,enterprise_id string,mobile_no string) (*models.Account, error)
	AddEnterprise(ctx context.Context, enterprise_email string, enterprise_username string, enterprise_pin string , enterprise_mobile string ) (*models.Account, error)
	GetEnterpriseUser(ctx context.Context)([]*models.Enterprise_User, error)
	UserofEnterprise(ctx context.Context , enterprise_id string ) ([]*models.Userof_Enterprise, error)
	DeleteEnterprise(ctx context.Context , enterprise_id string ) (*models.Account, error)
	DeleteUserofEnterprise(ctx context.Context , user_id string ) (*models.Account, error)
	UploadCsv(ctx context.Context , enterprise_id string ,filedata [][]string ) (*models.Account, error)
	Mod_enterprise(ctx context.Context ,enterprise_id string , enterprise_username string , enterprise_pin string ) (*models.Account, error)
	ShowOne_Enterprise(ctx context.Context , enterprise_id string ) (*models.ShowJoin_Enterprise, error)
	//QrcodeEnterprise(ctx context.Context , username string ,secret string) (*models.Account, error)
}

