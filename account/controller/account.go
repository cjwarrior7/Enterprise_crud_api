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

// GetOutboundRate - Route handler to fetch outbound rate
func (r *AccountController) Authenticate(c echo.Context) error {
	userName := c.QueryParam("userName")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	outboundCharge, _ := r.Usecase.AuthenticateUser(ctx,userName)
	return c.JSON(http.StatusOK, outboundCharge)
}

// NewRatesController - Initialize the controller object
func NewRatesController(e *echo.Echo, accoutnUsecase account.Usecase) {
	handler := &AccountController{
		Usecase: accoutnUsecase,
	}

	//zt/account/<account-id>/rates/outbound?toNumber<num>&fromNumber<num>
	e.GET("/v1/account/authentication/", handler.Authenticate)
}
