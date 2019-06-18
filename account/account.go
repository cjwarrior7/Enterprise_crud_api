package account


import (
	"context"
	"accountingService/models"
)

//Repository is repo layer for fetching Carrier domain objects
type Repository interface {
	GetByUsername(ctx context.Context, username string,secret string) (*models.Account, error)
	Enterprise_AddUser(ctx context.Context,name string,username string,user_pin string,user_email string,description string,enterprise_id string ,mobile_no string) (*models.Account, error)
	AddEnterpriseUser(ctx context.Context, enterprise_email string, enterprise_username string, enterprise_pin string, enterprise_mobile string) (*models.Account, error)
	GetAllEnterpriseUser(ctx context.Context)([]*models.Enterprise_User, error)
	GetUserofEnterprise(ctx context.Context,enterprise_id string)([]*models.Userof_Enterprise, error)
	DeleteEnterp(ctx context.Context , enterprise_id string)(*models.Account, error)
	DeleteUserofEnterp(ctx context.Context , enterprise_id string)(*models.Account, error)
	UploadCsvEnterp(ctx context.Context , enterprise_id string ,filedata [][] string)(*models.Account, error)
	Update_enterprise(ctx context.Context , enterprise_id string,enterprise_username string , enterprise_pin string)(*models.Account, error)
	ShowOne_Enterprise(ctx context.Context , enterprise_id string)(*models.ShowJoin_Enterprise, error)
}