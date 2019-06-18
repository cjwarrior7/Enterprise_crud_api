package models

// Account
type Account struct {
	Id int  `json:"id,omitempty"`
	Is_superuser  int `json:"is_superuser,omitempty"`
	Status int `json:"status"`
	Message string `json:"message,omitempty"`
	Rowsaffected string `json:"rowsaffected,omitempty"`
}
//type Register struct {
	//Status string `json:"status"`	
//	Message string `json:"message"`	
//}
type Enterprise_User struct {
	Email string `json:"enterprise_email"`	
	Username string `json:"enterprise_username"`	
	Pin string `json:"enterprise_pin"`	
	Mobile string `json:"enterprise_mobile"`	
}
type Userof_Enterprise struct {
	Name string `json:"name"`	
	Email string `json:"user_email"`	
	Username string `json:"username"`	
	Pin string `json:"user_pin"`	
	Mobile string `json:"user_mobile"`	
}


type ShowJoin_Enterprise struct {
	Is_superuser string `json:"is_superuser,omitempty"`	
	Is_active string `json:"is_active,omitempty"`
	Email string `json:"user_email,omitempty"`
	Username string `json:"username,omitempty"`	
	Pin string `json:"user_pin,omitempty"`	
	Mobile string `json:"user_mobile,omitempty"`
	Created string `json:"created_at,omitempty"`
	Updated string `json:"updated_at,omitempty"`
	Status string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`			
}

