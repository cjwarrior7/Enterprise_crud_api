package account

import ("fmt"
"context"
"net/http"
//"strconv"
"github.com/labstack/echo"
//"unicode"
"accountingService/account"
"accountingService/models"
"accountingService/logger"
"encoding/csv"
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
	secret := userDetails["secret"].(string)
	logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	fmt.Println("userName:", userName)
	fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.AuthenticateUser(ctx,userName, secret)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
//add enterprise
func (r *AccountController) Enterprise(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
    enterprise_email:= userDetails["enterprise_email"].(string)
	enterprise_username:=userDetails["enterprise_username"].(string)
	enterprise_pin:=userDetails["enterprise_pin"].(string)
	enterprise_mobile:=userDetails["enterprise_mobile"].(string)
	//subaccount_id:=userDetails["subaccount_id"].(string)
	logger.Logger.Info("AUTH Request Received with UserName:" + enterprise_username)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("username:",username)
	//fmt.Println("name:",name)
	//fmt.Println("pin_code:",pin_code)
   authResponse, _ := r.Usecase.AddEnterprise(ctx,enterprise_email , enterprise_username ,enterprise_pin , enterprise_mobile)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
// Register
func (r *AccountController) Register(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	username := userDetails["username"].(string)
	name:= userDetails["name"].(string)
	user_pin:=userDetails["user_pin"].(string)
	description:=userDetails["description"].(string)
	user_email:=userDetails["user_email"].(string)
	enterprise_id:=userDetails["enterprise_id"].(string)
	mobile_no:=userDetails["mobile_no"].(string)
	logger.Logger.Info("AUTH Request Received with UserName:" + username)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	fmt.Println("username:",username)
	fmt.Println("name:",name)
	fmt.Println("pin_code:",user_pin)

	authResponse, _ := r.Usecase.AddUser(ctx,name,username,user_pin,user_email,description,enterprise_id,mobile_no)
	if authResponse == nil {
		//result := make([]*models.Register, 0)
		//t := new(models.Register)
		// t.Status="1"
		// t.Message="account is not registered"
		// fmt.Println("hi:",t.Status)
		 // fmt.Println("hi:",t.Message)
		// result = append(result, t)
		
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
} 
//getEnterprise
func (r *AccountController) ShowEnterprise(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	
	//authId := c.Param("auth_id")
	//userName := userDetails["userName"].(string)
	//secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.GetEnterpriseUser(ctx)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
//show users of Enterprise

func (r *AccountController) ShowUserofEnterprise(c echo.Context ) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	enterprise_id := userDetails["enterprise_id"].(string)
	//secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.UserofEnterprise(ctx , enterprise_id )
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}

func (r *AccountController) DeleteEnp(c echo.Context ) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	enterprise_id := userDetails["enterprise_id"].(string)
	//secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.DeleteEnterprise(ctx , enterprise_id )
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}

func (r *AccountController) Upload(c echo.Context) error {
 //	var userDetails map[string]interface{}
     remodup:=[][]string{} //slice for removing duplicate from filedata
	//c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	//enterprise_id := userDetails["enterprise_id"].(string)
		// Read form fields
		enterprise_id := c.FormValue("enterprise_id")
		fmt.Println("enterprise_id",enterprise_id)

		//email := c.FormValue("email")
	
		//-----------
		// Read file
		//-----------
	
		// Source
		file, err := c.FormFile("file")
		fmt.Println("file ",file)

		if err != nil {
			fmt.Println("err ",err)
//			return err
		}
		src, err := file.Open()
		if err != nil {
			fmt.Println("err ",err)
//			return err
		}
		defer src.Close()
     // Destination
   /*  dst, err := os.Create(file.Filename)
       if err != nil {
		return err
				}
				defer dst.Close()
		
				// Copy
				if _, err = io.Copy(dst, src); err != nil {
					return err
				} */
		filedata ,err:=csv.NewReader(src).ReadAll()
		if err !=nil {
		 fmt.Println("error in filedata",err)
		 
		// panic(err)
		  return c.JSON(http.StatusOK,models.Account{Status:0,Message:"bad csv format"})
		  //return &models.Account{Status:0,string:"wrong csv format"}
	   }
	// for removing repeating element from filedata.
		for i := 0 ; i < len(filedata); i++ {
			// Scan slice for a previous element of the same value.
			exists := false
			for v := 0; v < i; v++ {
				if filedata[v][0] == filedata[i][0]{
					exists = true
					break
				}
			}
			// If no previous element exists, append this one.
			if !exists {
				remodup = append(remodup , filedata[i][:])
				fmt.Println("removedup :",remodup)
			}
		}
		fmt.Println("removingduplicate:",remodup)
	
	  // fmt.Println("hi:filedata:",filedata)
	  //removing 
     
		
		

	 


	//secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.UploadCsv(ctx,enterprise_id,remodup) 
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
func (r *AccountController) DeleteUser(c echo.Context ) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
	user_id := userDetails["user_id"].(string)
	//secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.DeleteUserofEnterprise(ctx , user_id )
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}









/*func (r *AccountController) QrcodeEnp(c echo.Context ) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
}
	user_id := userDetails["userName"].(string)
	secret := userDetails["secret"].(string)
	//logger.Logger.Info("AUTH Request Received with UserName:" + userName)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("userName:", userName)
	//fmt.Println("secret:", secret)

	authResponse, _ := r.Usecase.QrcodeEnterprise(ctx,user_id,secret)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
*/

// modify enterprise
func (r *AccountController) Modify_enterprise(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
    enterprise_id:= userDetails["enterprise_id"].(string)
	enterprise_username:=userDetails["enterprise_username"].(string)
	enterprise_pin:=userDetails["enterprise_pin"].(string)
	//enterprise_mobile:=userDetails["enterprise_mobile"].(string)
	//subaccount_id:=userDetails["subaccount_id"].(string)
	logger.Logger.Info("AUTH Request Received with UserName:" + enterprise_username)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("username:",username)
	//fmt.Println("name:",name)
	//fmt.Println("pin_code:",pin_code)
   authResponse, _ := r.Usecase.Mod_enterprise(ctx ,enterprise_id,enterprise_username,enterprise_pin )
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
} 

// ShowOneEnterprise with join 

func (r *AccountController) ShowOneEnterprise(c echo.Context) error {
	var userDetails map[string]interface{}
	c.Bind(&userDetails)
	//authId := c.Param("auth_id")
    enterprise_id:= userDetails["enterprise_id"].(string)
	//enterprise_username:=userDetails["enterprise_username"].(string)
	//enterprise_pin:=userDetails["enterprise_pin"].(string)
	//enterprise_mobile:=userDetails["enterprise_mobile"].(string)
	//subaccount_id:=userDetails["subaccount_id"].(string)
	// logger.Logger.Info("AUTH Request Received with UserName:" + enterprise_id)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//fmt.Println("username:",username)
	//fmt.Println("name:",name)
	//fmt.Println("pin_code:",pin_code)
   authResponse, _ := r.Usecase.ShowOne_Enterprise(ctx,enterprise_id)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
} 






// NewRatesController - Initialize the controller object
func NewAccountController(e *echo.Echo, accoutnUsecase account.Usecase) {
	handler := &AccountController{
		Usecase: accoutnUsecase,
	}

	//zt/account/<account-id>/rates/outbound?toNumber<num>&fromNumber<num>
	e.POST("/v1/account/authentication/", handler.Authenticate)
	e.POST("/v1/account/addEnterprise/", handler.Enterprise)
	e.POST("/v1/account/AddUser/", handler.Register)
	e.POST("/v1/account/Upload/", handler.Upload)
	e.DELETE("/v1/account/DeleteEnterprise/", handler.DeleteEnp)
	e.DELETE("/v1/account/DeleteUser/", handler.DeleteUser)
	e.GET("/v1/account/ShowEnterprise/", handler.ShowEnterprise)
	e.GET("/v1/account/ShowUsersofEnterprise/", handler.ShowUserofEnterprise)
	e.PUT("/v1/account/Modify_enterprise/", handler.Modify_enterprise)
	e.GET("/v1/account/ShowOneEnterprise/", handler.ShowOneEnterprise)
	// e.POST("/v1/account/qrcode/", handler.qrcodegenrator)

}
