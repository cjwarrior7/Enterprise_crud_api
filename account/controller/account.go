package account

import (
"context"
"net/http"

"github.com/labstack/echo"

"accountingService/account"

)

// RatesController - Controller for rating
type AccountController struct {
	Usecase account.Usecase
}

// Authenticate
func (r *AccountController) Authenticate(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	userName := userDetails["userName"].(string)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authResponse, _ := r.Usecase.AuthenticateUser(ctx,userName)
	return c.JSON(http.StatusOK, authResponse)
}

// NewRatesController - Initialize the controller object
func NewAccountController(e *echo.Echo, accoutnUsecase account.Usecase) {
	handler := &AccountController{
		Usecase: accoutnUsecase,
	}

	//zt/account/<account-id>/rates/outbound?toNumber<num>&fromNumber<num>
	e.POST("/v1/account/authentication/", handler.Authenticate)
}
