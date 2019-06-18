package repository

import (
	//"io/ioutil"
	//"crypto/md5"
	"context"
	"database/sql"
	"fmt"
	"accountingService/logger"
	"accountingService/models"
	"accountingService/account"
	_"github.com/lib/pq"
	"net/smtp"
	"strconv"
	

//	"io"
)

type accountRepository struct {
	DbConn *sql.DB
}

// NewCarrierRepository - Repository for carrier
func NewAccountRepository(Conn *sql.DB)  account.Repository{
	return &accountRepository{
     DbConn: Conn,
	}
}

func  (c *accountRepository) GetByUsername(ctx context.Context, username string, secret string) (*models.Account, error) {

//	hashMd5 := md5.New()
//	io.WriteString(hashMd5, secret)
//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))

fmt.Println("hello username repository", username)
fmt.Println("hello secret repository",secret)


	query := fmt.Sprintf("SELECT id,is_superuser FROM account_login WHERE is_active=true and username = '%s' " +
		"and password = '%s'", username,secret)// secret md5Hash)
		fmt.Printf("hi:%s", query)
	rows, err := c.DbConn.QueryContext(ctx, query)
	//query:=fmt.Sprintf("insert into add_users(subaccount_id) values('%s')",subaccount_id )
	//rows, err = c.DbConn.ExecContext(ctx, query)

	fmt.Println("rows:hi", rows)
	fmt.Println("err :hi", err)
	


	if err != nil {
		result := make([]*models.Account, 0)
		t := new(models.Account)
		t.Id=0
		t.Is_superuser=0
		t.Status=200
		t.Message="account information not saved"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
		result = append(result, t)
        fmt.Println("hi:result",result[0])

		logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching account_details from database")
		//return nil, err
		return result[0], nil
	}
	logger.Logger.Info("In accountRepo:GetByUsername")
	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()
          
	result := make([]*models.Account, 0)
	for rows.Next() {
		t := new(models.Account)
		err := rows.Scan(&t.Id,&t.Is_superuser)
	    t.Status=1
//		t.Message="success"
//		fmt.Println("hi:",t.Id)
//		fmt.Println("hi:",t.Is_superuser)
		//fmt.Println("hi:",t.Status)
		//fmt.Println("hi:",t.Message)
		if err != nil {
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
		return nil, err
			
		}
		result = append(result, t )
	
	
		
		
		//result = append(result, t)


		//insert(t.Id int)
		  //subaccount_id:=t.Id
		 // superuser:=t.Is_superuser

		 // fmt.Println("accountid: ",subaccount_id)

	      //subaccount_id:=(string)(account_id)
		 // fmt.Println("hi chandresh:",subaccount_id)
		  //query:=fmt.Sprintf("Insert into add_users(subaccount_id,superuser) values('%d','%d')",subaccount_id,superuser);
		 //_, err = c.DbConn.ExecContext(ctx, query)
		 //fmt.Println("err:",err)
		 fmt.Println("hi result:",result[0])
	}
      
	//fmt.Println("result",result[0])

	if len(result) == 0 {
	    logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
		//return nil, fmt.Errorf("carrier not found")
		//contents:=models.Account{ Message:"failed",Status:"200"}
		return &models.Account{ Message:"wrong username or password",Status:200}, nil
	}
	return result[0], nil
}
//add enterprise
func  (c *accountRepository) AddEnterpriseUser(ctx context.Context, enterprise_email string, enterprise_username string,enterprise_pin string,enterprise_mobile string) (*models.Account, error) {

	//	hashMd5 := md5.New()
	//	io.WriteString(hashMd5, secret)
	//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))

//check if email already exists
	query := fmt.Sprintf("select enterprise_email from enterprise where enterprise_email='%s' ",enterprise_email)
	rows, err := c.DbConn.QueryContext(ctx, query)
	fmt.Println("rows:%s",rows)
	fmt.Println("err:%s",err)
	

	if err != nil {
        logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching account_details from database")
//		return result[0], nil
		return &models.Account{Status:1,Message:"Failed"}, err
	}

//add defer function
defer func() {
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}()




	result := make([]*models.Account, 0)
	for rows.Next() {

		t := new(models.Account)
/*
		t.Status=0
		t.Message="email_id already exist"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
*/
		err := rows.Scan(&t.Message)
		fmt.Println("err while scan:%s",err)
		fmt.Println(t.Message)
		if err != nil {
			fmt.Println("here rows.scan is return errors")
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
//			return nil, err
            //return &models.Account{Status:0,Message:"email_id already exists"}, err
		}
		result = append(result, t)		
	} 

	
	if len(result) != 0 {
		fmt.Println("here length  is not zero ")	
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
        return &models.Account{Status:0,Message:"Email ID already exists"}, err

	} else {

// if email does not exist
//now check if mobile_no. already exists
query := fmt.Sprintf("select enterprise_mobile from enterprise where enterprise_mobile='%s' ",enterprise_mobile)
	rows, err := c.DbConn.QueryContext(ctx, query)
	fmt.Println("rows:%s",rows)
	fmt.Println("err:%s",err)
	if err != nil {
        logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching account_details from database")
//		return result[0], nil
		return &models.Account{Status:0,Message:"Failed"}, err
	}

//add defer function
defer func() {
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}()




	result := make([]*models.Account, 0)
	for rows.Next() {

		t := new(models.Account)
/*
		t.Status=0
		t.Message="email_id already exist"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
*/
		err := rows.Scan(&t.Message)
		fmt.Println("err while scan:%s",err)
		if err != nil {
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
			return nil, err
            //return &models.Account{Status:0,Message:"mobile_no. already exists"}, err
		}
		result = append(result, t)		
	} 

	
	if len(result) != 0 {
		fmt.Println("here length result is not zero we are in mobile check")	
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
        return &models.Account{Status:0,Message:"mobile_no already exists"}, err

	} else{
		query := fmt.Sprintf("select enterprise_username from enterprise where enterprise_username='%s' ",enterprise_username)
	rows, err := c.DbConn.QueryContext(ctx, query)
	fmt.Println("rows:%s",rows)
	fmt.Println("err:%s",err)
	if err != nil {
        logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching account_details from database")
//		return result[0], nil
		return &models.Account{Status:0,Message:"Failed"}, err
	}

//add defer function
defer func() {
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}()




	result := make([]*models.Account, 0)
	for rows.Next() {

		t := new(models.Account)
/*
		t.Status=0
		t.Message="email_id already exist"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
*/
		
        //fmt.Println("hi:%s",t.Message)
		err := rows.Scan(&t.Message)
		fmt.Println("hi:",t.Message)
		fmt.Println("err while scan:",err)
		if err != nil {
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
			return nil, err
            //return &models.Account{Status:0,Message:"mobile_no. already exists"}, err
		}
		result = append(result, t)		
	} 

	
	if len(result) != 0 {
		fmt.Println("here length result is not zero we are in username check")	
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
        return &models.Account{Status:0,Message:"enterprise_username already exists"}, err

	} else{
		


// it will pass check of username ,mobile no,  email_id
  id:=0
  fmt.Println("hello username repository", enterprise_username)
	query = fmt.Sprintf("INSERT into enterprise(enterprise_username,enterprise_email,enterprise_mobile,enterprise_pin) values('%s','%s','%s','%s') RETURNING enterprise_id ",enterprise_username,enterprise_email,enterprise_mobile,enterprise_pin)
			fmt.Printf("hi:%s", query)
			//fmt.Printf("hi:%d", id)
		rows, err := c.DbConn.QueryContext(ctx, query)
		
		if err != nil{
			result := make([]*models.Account, 0)
			t := new(models.Account)
			t.Status=0
			t.Message="account information not saved"
			fmt.Println("hi:",t.Status)
			fmt.Println("hi:",t.Message)
			result = append(result, t)
			fmt.Println("hi:result",result[0])
			//return result[0], nil 
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching account_details from database")
			return result[0], nil
		}
		// i am adding here today
//add defer function
defer func() {
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}()




	result := make([]*models.Account, 0)
	for rows.Next() {

		t := new(models.Account)
/*
		t.Status=0
		t.Message="email_id already exist"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
*/
        fmt.Println("hi:",id)
		err := rows.Scan(&id)
		fmt.Println("hi:",id)
		fmt.Println("err while scan:",err)
		if err != nil {
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
			return nil, err
            //return &models.Account{Status:0,Message:"mobile_no. already exists"}, err
		}
		result = append(result, t)		
	} 


		query = fmt.Sprintf("INSERT into account_login(is_superuser,username,password,is_active,enterprise_id) values('%d','%s','%s',true,'%d')",2,enterprise_username,enterprise_pin,id)//md5Hash)
			fmt.Printf("hi:%s", query)
		rows, err = c.DbConn.QueryContext(ctx,query)
		//fmt.Println("rows:hi", rows)
		//fmt.Println("err :hi", err)
		//send username and enterprise_pin to enterprise
	if err==nil {
		hostURL:="smtp.gmail.com"
		hostPort:="587"
		emailSender:="smartoffice591@gmail.com"
		password:="msdhoni7"
		emailReceiver:=enterprise_email
		emailAuth :=smtp.PlainAuth(
			"",
			emailSender,
			password,
			hostURL,
		)
		
	   // credentials:=fmt.Sprintf("%s%s%s%s","username:", username ,"pin_code:",pin_code)
		msg:="To: " + emailReceiver + "\r\n" +
		"Subject : send by admin \n\n"+
		// "\r\n username:" +username+ "\r\n pin:" +pin_code+ "\r\n"
		"your username is : "+enterprise_username+ " and pin is : "+enterprise_pin+"\n\n" 
   
		err=smtp.SendMail(hostURL + ":" + hostPort,
		 emailAuth,
		 emailSender, 
		 []string{emailReceiver},
		 []byte (msg))
		if err !=nil {
			fmt.Print("Error:",err)
		} 
		fmt.Print("Email sent successfully")

	   }

		if err != nil {
			result := make([]*models.Account, 0)
			t := new(models.Account)
			t.Status=0
			t.Message="account information not saved"
			fmt.Println("hi:",t.Status)
			fmt.Println("hi:",t.Message)
			result = append(result, t)
			fmt.Println("hi:result",result[0])
			//return result[0], nil 
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching account_details from database")
			return result[0], nil
		}
		logger.Logger.Info("In accountRepo:RegisterForPin")
		defer func() {
			//err := rows.Close()
			if err != nil {	
				panic(err)
			}
		}()
		
	   result = make([]*models.Account, 0)
		if rows!=nil {
			t := new(models.Account)
			t.Status=1
			t.Message="account successfully registered"
			fmt.Println("hi:",t.Status)
			fmt.Println("hi:",t.Message)
			if err != nil {
				logger.Logger.WithError(err).WithField("query", query).
				Errorf("Error while fetching carriers from database")
				return nil, err
			}
		    result = append(result, t)
		} 
	
		
			if len(result) == 0 {
			
				logger.Logger.WithError(err).WithField("query", query).
				Errorf("Error while fetching carriers from database")
				return nil, fmt.Errorf("carrier not found")
			}
			return result[0], nil
		}  //else inside inner else closing

	}	//inner else closing	
		
	}  //outer else closing

}//function ending




//add user by enterprise
func (c *accountRepository) Enterprise_AddUser(ctx context.Context,name string,username string,user_pin string,user_email string,description string,enterprise_id string ,mobile_no string) (*models.Account, error)  {
   
/*	if (user_email==""){
		result := make([]*models.Account, 0)
		t := new(models.Account)
		t.Status=2
		t.Message="please enter email_id"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
		result = append(result, t)
        return result[0], nil 	
} 
*/
//runs without error start
//check if email already exists
query := fmt.Sprintf("select user_email from enterprise_addusers where user_email='%s' ",user_email)
rows, err := c.DbConn.QueryContext(ctx, query)
fmt.Println("rows:%s",rows)
fmt.Println("err:%s",err)
if err != nil {
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching account_details from database")
//		return result[0], nil
	return &models.Account{Status:1,Message:"Failed"}, err
}

//add defer function
defer func() {
err := rows.Close()
if err != nil {
	panic(err)
}
}()




result := make([]*models.Account, 0)
for rows.Next() {

	t := new(models.Account)
/*
	t.Status=0
	t.Message="email_id already exist"
	fmt.Println("hi:",t.Status)
	fmt.Println("hi:",t.Message)
*/
	err := rows.Scan(&t.Message)
	fmt.Println("err while scan:%s",err)
	fmt.Println(t.Message)
	if err != nil {
		fmt.Println("here rows.scan is return errors")
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//			return nil, err
		//return &models.Account{Status:0,Message:"email_id already exists"}, err
	}
	result = append(result, t)		
} 


if len(result) != 0 {
	fmt.Println("here length  is not zero ")	
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
	return &models.Account{Status:0,Message:"Email ID already exists"}, err

} else {

// if email does not exist
//now check if mobile_no. already exists
query := fmt.Sprintf("select user_mobile from enterprise_addusers where user_mobile='%s' ",mobile_no)
rows, err := c.DbConn.QueryContext(ctx, query)
fmt.Println("rows:%s",rows)



//add defer function
defer func() {
err := rows.Close()
if err != nil {
	panic(err)
}
}()




result := make([]*models.Account, 0)
for rows.Next() {

	t := new(models.Account)
/*
	t.Status=0
	t.Message="email_id already exist"
	fmt.Println("hi:",t.Status)
	fmt.Println("hi:",t.Message)
*/
	err := rows.Scan(&t.Message)
	fmt.Println("err while scan:%s",err)
	if err != nil {
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
		return nil, err
		//return &models.Account{Status:0,Message:"mobile_no. already exists"}, err
	}
	result = append(result, t)		
} 


if len(result) != 0 {
	fmt.Println("here length result is not zero we are in mobile check")	
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
	return &models.Account{Status:0,Message:"mobile_no already exists"}, err

} 	else{


    //runs without error end
	fmt.Println("hello username repository", username)
     fmt.Println("hello PIN repository",user_pin)
    //conv,_:=strconv.Atoi(subaccount_id)
	query := fmt.Sprintf("Insert into enterprise_addusers(enterprise_id,name,username,user_pin,user_email,description,user_mobile) values('%s','%s','%s','%s','%s','%s','%s')",enterprise_id ,name,username,user_pin,user_email,description,mobile_no)
		fmt.Printf("hi:%s", query)
	rows , err := c.DbConn.ExecContext(ctx, query)
    fmt.Print("error: %s",err)



	//send username and pincode
	if err==nil {
	 hostURL:="smtp.gmail.com"
	 hostPort:="587"
	 emailSender:="smartoffice591@gmail.com"
	 password:="msdhoni7"
	 emailReceiver:=user_email
	 emailAuth :=smtp.PlainAuth(
		 "",
		 emailSender,
		 password,
		 hostURL,
	 )
	 
	// credentials:=fmt.Sprintf("%s%s%s%s","username:", username ,"pin_code:",pin_code)
	 msg:="To: " + emailReceiver + "\r\n" +
	 "Subject : sendbyadmin \n\n"+
	 // "\r\n username:" +username+ "\r\n pin:" +pin_code+ "\r\n"
	 "your username is : "+username+ " and pin is : "+user_pin+"\n\n"

	 err=smtp.SendMail(hostURL + ":" + hostPort,
	  emailAuth,
	  emailSender, 
	  []string{emailReceiver},
	  []byte (msg))
	 if err !=nil {
		 fmt.Print("Error:",err)
	 } 
	 fmt.Print("Email sent successfully")
	}
	//email sent code end
	//if(rows!=nil){
	//fmt.Println("hi rows:",rows)
	// }

	if err != nil {
	    result := make([]*models.Account, 0)
		t := new(models.Account)
		t.Status=0
		t.Message="account information not saved"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
		result = append(result, t)
        fmt.Println("hi:result",result[0])
		//return result[0], nil 
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching account_details from database")
		return result[0], nil
	}
	logger.Logger.Info("In accountRepo:RegisterForPin")
	defer func() {
		//err := rows.Close()
		if err != nil {	
			panic(err)
		}
	}()
	
   result := make([]*models.Account, 0)
	if rows!=nil {
		t := new(models.Account)
		t.Status=1
		t.Message="account successfully registered by enterprise"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
		if err != nil {
            logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
			return nil, err
		}
	result = append(result, t)
	} 
       if len(result) == 0 {
	    logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
		//return nil, fmt.Errorf("carrier not found")
		return &models.Account{Status:0,Message:"Failed"},nil
		}
		return result[0], nil
	
	}//Inner else closing
} //outer else closing
} //function closing

//GetEnterprise
func  (c *accountRepository) GetAllEnterpriseUser(ctx context.Context) ([]*models.Enterprise_User, error) {

	//	hashMd5 := md5.New()
	//	io.WriteString(hashMd5, secret)
	//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))
	
	//fmt.Println("hello username repository", enterprise_username)
	query := fmt.Sprintf("select enterprise_email,enterprise_username, enterprise_pin, enterprise_mobile from enterprise order by enterprise_id DESC")//md5Hash)
			fmt.Printf("hi:%s", query)
			rows, err := c.DbConn.QueryContext(ctx, query)
			if err != nil {
			   logger.Logger.WithError(err).WithField("query", query).
				  Errorf("Error while fetching project_details By ID from database")
			   fmt.Println("repository/account.go error\n", err)
			   fmt.Println("repository/account.go rows\n", rows)
			   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
			   return nil, err
			}
	  
			if rows == nil {
			   fmt.Println("rows are nil\n", err)
			}
	  
			logger.Logger.Info("In accountRepo:GetProjectById")
			defer func() {
			   err := rows.Close()
			   if err != nil {
				  panic(err)
			   }
			}()
	  
			fmt.Printf("repository/account.go\n")
	  
			result := make([]*models.Enterprise_User, 0)
			for rows.Next() {
			   t := new(models.Enterprise_User)
			   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
			   err := rows.Scan(&t.Email, &t.Username, &t.Pin, &t.Mobile)
	  
			   if err != nil {
				  logger.Logger.WithError(err).WithField("query", query).
					 Errorf("Error while fetching carriers from database")
				 // fmt.Printf("rows.next giving error\n", err)
				  return nil, err
			   }
			   result = append(result, t)
			}
	  
			fmt.Println("result ", result)
	  
			if len(result) == 0 {
			   logger.Logger.WithError(err).WithField("query", query).
				  Errorf("Error while fetching carriers from database")
			   fmt.Println("len(result) is zero\n", err)
	  
			  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
			   //fmt.Printf("status \n", status)
			   //result = append(result, status)
			   return result, err
			   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
	  
			}
		
			return result, err
	}


	//get user of enterprise
	func  (c *accountRepository) GetUserofEnterprise(ctx context.Context , enterprise_id string) ([]*models.Userof_Enterprise , error) {

		//	hashMd5 := md5.New()
		//	io.WriteString(hashMd5, secret)
		//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))
		
		//fmt.Println("hello username repository", enterprise_username)
		query := fmt.Sprintf("select name , user_email,username, user_pin, user_mobile from enterprise_addusers where enterprise_id='%s' order by user_id Desc",enterprise_id)//md5Hash)
				fmt.Printf("hi:%s", query)
				rows, err := c.DbConn.QueryContext(ctx, query)
				if err != nil {
				   logger.Logger.WithError(err).WithField("query", query).
					  Errorf("Error while fetching project_details By ID from database")
				   fmt.Println("repository/account.go error\n", err)
				   fmt.Println("repository/account.go rows\n", rows)
				   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
				   return nil, err
				}
		  
				if rows == nil {
				   fmt.Println("rows are nil\n", err)
				}
		  
				logger.Logger.Info("In accountRepo:GetProjectById")
				defer func() {
				   err := rows.Close()
				   if err != nil {
					  panic(err)
				   }
				}()
		  
				fmt.Printf("repository/account.go\n")
		  
				result := make([]*models.Userof_Enterprise, 0)
				for rows.Next() {
				   t := new(models.Userof_Enterprise)
				   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
				   err := rows.Scan(&t.Name, &t.Email, &t.Username, &t.Pin, &t.Mobile)
		  
				   if err != nil {
					  logger.Logger.WithError(err).WithField("query", query).
						 Errorf("Error while fetching carriers from database")
					 // fmt.Printf("rows.next giving error\n", err)
					  return nil, err
				   }
				   result = append(result, t)
				}
		  
				fmt.Println("result ", result)
		  
				if len(result) == 0 {
				   logger.Logger.WithError(err).WithField("query", query).
					  Errorf("Error while fetching carriers from database")
				   fmt.Println("len(result) is zero\n", err)
		  
				  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
				   //fmt.Printf("status \n", status)
				   //result = append(result, status)
				   return result, err
				   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
		  
				}
			
				return result, err
		}
		//delete enterprise
func  (c *accountRepository) DeleteEnterp(ctx context.Context,enterprise_id string) (*models.Account, error) {

			//	hashMd5 := md5.New()
			//	io.WriteString(hashMd5, secret)
			//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))
			
			//fmt.Println("hello username repository", enterprise_username)
query := fmt.Sprintf("delete from enterprise where enterprise_id=%s",enterprise_id)//md5Hash)
					fmt.Printf("hi:%s", query)
					rows, err := c.DbConn.ExecContext(ctx, query)
					rowsUpdated,err:=rows.RowsAffected()
					fmt.Println("rows:",rows)
					if err != nil {
						logger.Logger.WithError(err).WithField("query", query).
						   Errorf("Error while fetching project_details By ID from database")
						fmt.Println("repository/account.go error\n", err)
						fmt.Println("repository/account.go rows\n", rows)
					  return &models.Account{Status:0,Message:"Failed"}, nil
						//return nil, err
					 }

query = fmt.Sprintf("delete from account_login where enterprise_id=%s",enterprise_id)//md5Hash)
                    fmt.Printf("hi:%s", query)
					rows, err = c.DbConn.ExecContext(ctx, query)
					rowsUpdated1,err:=rows.RowsAffected()
		            fmt.Println("rows:",rows)
					if err != nil {
					   logger.Logger.WithError(err).WithField("query", query).
						  Errorf("Error while fetching project_details By ID from database")
					   fmt.Println("repository/account.go error\n", err)
					   fmt.Println("repository/account.go rows\n", rows)
					 return &models.Account{Status:0,Message:"Failed"}, nil
					   //return nil, err
					}
			  
					if rows == nil {
					   fmt.Println("rows are nil\n", err)
					}
			  //comment by me
					/*logger.Logger.Info("In accountRepo:GetProjectById")
					defer func() {
					   err := rows.Close()
					   if err != nil {
						  panic(err)
					   }
					}()
			  
					fmt.Printf("repository/account.go\n")
			  
					result := make([]*models.Account, 0)
					for rows.Next() {
					   t := new(models.Account)
					   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
					   err := rows.Scan(&t.Message,&t.Status)
					   fmt.Printf("Message:%s ",t.Message)
					   fmt.Printf("status:%s ",t.Status)
					   if err != nil {
						  logger.Logger.WithError(err).WithField("query", query).
							 Errorf("Error while fetching carriers from database")
						 // fmt.Printf("rows.next giving error\n", err)
						  return &models.Account{Status:0,Message:"Failed"}, nil
					   }
					   result = append(result, t)
					}*/
			  
					//fmt.Println("result ", result)
					//comment by me end
			  
					if ( rowsUpdated == 0 || rowsUpdated1 == 0 ) {
						fmt.Println("no result selected therefore length of result is zero")
					   logger.Logger.WithError(err).WithField("query", query).
						  Errorf("Error while fetching carriers from database")
					   fmt.Println("len(result) is zero\n", err)
			  
					  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
					   //fmt.Printf("status \n", status)
					   //result = append(result, status)
					   
					    return &models.Account{Status:0,Message:"Failed"}, nil
			  
					} 
					
				
					return &models.Account{Status:1,Message:"deleted successfully"}, nil
			}
			//delete user of enterprise
func  (c *accountRepository) DeleteUserofEnterp(ctx context.Context,user_id string) (*models.Account, error) {

				//	hashMd5 := md5.New()
				//	io.WriteString(hashMd5, secret)
				//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))
				
				//fmt.Println("hello username repository", enterprise_username)
				query := fmt.Sprintf("delete from  enterprise_addusers where user_id=%s",user_id)//md5Hash)
						fmt.Printf("hi:%s", query)
						rows, err := c.DbConn.ExecContext(ctx, query)
						rowsUpdated,err:=rows.RowsAffected()
						if err != nil {
						   logger.Logger.WithError(err).WithField("query", query).
							  Errorf("Error while fetching project_details By ID from database")
						   fmt.Println("repository/account.go error\n", err)
						   fmt.Println("repository/account.go rows\n", rows)
						   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
						   //return nil, err
						   return &models.Account{Status:0,Message:"Failed"}, nil
						}
				  
						if rows == nil {
						   fmt.Println("rows are nil\n", err)
						}

				  //comment by me
						/*logger.Logger.Info("In accountRepo:GetProjectById")
						defer func() {
						   err := rows.Close()
						   if err != nil {
							  panic(err)
						   }
						}()
				  
						fmt.Printf("repository/account.go\n")
				  
						result := make([]*models.Account, 0)
						for rows.Next() {
						   t := new(models.Account)
						   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
						   err := rows.Scan(&t.Message,&t.Status)
				            fmt.Printf("%s:",t.Message)
						   if err != nil {
							  logger.Logger.WithError(err).WithField("query", query).
								 Errorf("Error while fetching carriers from database")
							 // fmt.Printf("rows.next giving error\n", err)
							  return nil, err
						   }
						   result = append(result, t)
						}
				  
						fmt.Println("result ", result)*/
						//comment by me
				  
						if (rowsUpdated==0){
						   logger.Logger.WithError(err).WithField("query", query).
							  Errorf("Error while fetching carriers from database")
						   fmt.Println("len(result) is zero\n", err)
				  
						  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
						   //fmt.Printf("status \n", status)
						   //result = append(result, status)
						   return &models.Account{Status:0,Message:"Failed"}, nil
						   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
				  
						}
					
						return &models.Account{Status:1,Message:"deleted successfully"}, nil
				}
func  (c *accountRepository) UploadCsvEnterp(ctx context.Context,enterprise_id string,filedata [][]string) (*models.Account, error) {
	result := make([]*models.Account, 0) 
		// var grid [10][4]int
		  // fmt.Println(filedata[:][i])
		 // for i := range filedata {
		  fmt.Println("0:",filedata[0][0])
		  fmt.Println("1:",filedata[0][1])

		  //fmt.Println("1:",filedata[0][1])
		  //fmt.Println("2:",filedata[0][2])
		 // }func IsLetter(s string) bool {
  /* func IsLetter(s string) bool {
    for _, r := range filedata {
 if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
	return false
		}
				}
				return true
			} */

	  for i :=0; i<len(filedata) ; i++ {
		  
		   if(filedata[i][0]==" "){
			   continue
			   //return &models.Account{Status:0,Message:"failed rollnum on certain row is empty"},nil
			  } else {
				  
				 
				  if _,err:=strconv.Atoi(filedata[i][0]); err!=nil {
					 fmt.Println(":err",err)
					 continue
				     
				  }
				  if(filedata[i][0][0]=='9' && filedata[i][0][1]=='1') {
				   if(len(filedata[i][0])<12 || len(filedata[i][0])>14 ){
					fmt.Println(len(filedata[i][0]))
				     continue 

					 //
					//continue outer
					//return &models.Account{Status:0,Message:"Mobile number invalid"},nil
				   }

				  // fmt.Println("filedata[i][2] ", filedata[i][2])

				
				
				  
			query:= fmt.Sprintf("Insert into uploadfile(enterprise_id , mobile_no , name , city , state , pincode , country) values('%s' , '%s' ,'%s' ,'%s','%s','%s','%s')", enterprise_id , filedata[i][0], filedata[i][1], filedata[i][2],filedata[i][3],filedata[i][4],filedata[i][5])//md5Hash)
            fmt.Printf("hi:%s", query)
			rows, err := c.DbConn.ExecContext(ctx, query)
			//rowsUpdated,err:=rows.RowsAffected()
			if err != nil {
			   logger.Logger.WithError(err).WithField("query", query).
				  Errorf("Error while fetching project_details By ID from database")
			   fmt.Println("repository/account.go error\n", err)
			   fmt.Println("repository/account.go rows\n", rows)
			   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
			   //return nil, err
			   return &models.Account{Status:0,Message:"Failed"}, nil
			}
			//result := make([]*models.Account, 0)
			if rows!=nil {
				t := new(models.Account)
				t.Status=1
				t.Message="account successfully registered"
				fmt.Println("hi:",t.Status)
				fmt.Println("hi:",t.Message)
				if err != nil {
					logger.Logger.WithError(err).WithField("query", query).
					Errorf("Error while fetching carriers from database")
					return nil, err
				}
				result = append(result, t)
			}  
				 
		}//if terminate			 		 
          }//else terminate
		}//for terminate
	
		  // fmt.Println(filedata[i][0],filedata[i][1],filedata[i][2])
	
		   // ins := "INSERT INTO upcsv(rollnum,name) VALUES($1,$2)"
		  // tags := filedata[:][i]
		  // _,err = db.Exec(ins,pq.Array(&tags))//
		  
			//if err != nil {
			//panic(err)
	
			//}
/*query := fmt.Sprintf("Insert into uploadcsv(enterprise_id , rollnum , name , age) values('%d' , '%d' ,'%s' ,'%d')",enterprise_id )//md5Hash)
							fmt.Printf("hi:%s", query)
							rows, err := c.DbConn.ExecContext(ctx, query)
							rowsUpdated,err:=rows.RowsAffected()
							if err != nil {
							   logger.Logger.WithError(err).WithField("query", query).
								  Errorf("Error while fetching project_details By ID from database")
							   fmt.Println("repository/account.go error\n", err)
							   fmt.Println("repository/account.go rows\n", rows)
							   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
							   //return nil, err
							   return &models.Account{Status:0,Message:"Failed"}, nil

							   
							}
					  
							if rows == nil {
							   fmt.Println("rows are nil\n", err)
							}
	
					  //comment by me
							/*logger.Logger.Info("In accountRepo:GetProjectById")
							defer func() {
							   err := rows.Close()
							   if err != nil {
								  panic(err)
							   }
							}()
					  
							fmt.Printf("repository/account.go\n")
					  
							result := make([]*models.Account, 0)
							for rows.Next() {
							   t := new(models.Account)
							   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
							   err := rows.Scan(&t.Message,&t.Status)
								fmt.Printf("%s:",t.Message)
							   if err != nil {
								  logger.Logger.WithError(err).WithField("query", query).
									 Errorf("Error while fetching carriers from database")
								 // fmt.Printf("rows.next giving error\n", err)
								  return nil, err
							   }
							   result = append(result, t)
							}
					  
							fmt.Println("result ", result)*/
							//comment by me
					  
							/*if (rowsUpdated==0){
							   logger.Logger.WithError(err).WithField("query", query).
								  Errorf("Error while fetching carriers from database")
							   fmt.Println("len(result) is zero\n", err)
					  
							  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
							   //fmt.Printf("status \n", status)
							   //result = append(result, status)
							   return &models.Account{Status:0,Message:"Failed"}, nil
							   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
					  
							}*/
							if len(result) == 0 {
							//	logger.Logger.WithError(err).WithField("query", query).
							//	Errorf("Error while fetching carriers from database")
								//fmt.Println("len(result) is zero\n", err)
					   
							   // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
								//fmt.Printf("status \n", status)
								//result = append(result, status)
								fmt.Println("nothing is in result")
								return &models.Account{Status:0,Message:"failed"},nil
								// return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
					   
							 }
						
							return &models.Account{Status:1,Message:"csv inserted"}, nil 
						}
// update_enterprise

func  (c *accountRepository) Update_enterprise(ctx context.Context, enterprise_id string ,enterprise_username string,enterprise_pin string) (*models.Account, error) {
	query := fmt.Sprintf("select enterprise_username from enterprise where enterprise_username='%s' ",enterprise_username)
	rows, err := c.DbConn.QueryContext(ctx, query)
	fmt.Println("rows:%s",rows)
	fmt.Println("err:%s",err)
	if err != nil {
        logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching account_details from database")
//		return result[0], nil
		return &models.Account{Status:0,Message:"Failed"}, err
	}

//add defer function
defer func() {
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}()




	result := make([]*models.Account, 0)
	for rows.Next() {

		t := new(models.Account)
/*
		t.Status=0
		t.Message="email_id already exist"
		fmt.Println("hi:",t.Status)
		fmt.Println("hi:",t.Message)
*/
		
        //fmt.Println("hi:%s",t.Message)
		err := rows.Scan(&t.Message)
		fmt.Println("hi:",t.Message)
		fmt.Println("err while scan:",err)
		if err != nil {
			logger.Logger.WithError(err).WithField("query", query).
			Errorf("Error while fetching carriers from database")
			return nil, err
            //return &models.Account{Status:0,Message:"mobile_no. already exists"}, err
		}
		result = append(result, t)		
	} 

	
	if len(result) != 0 {
		fmt.Println("here length result is not zero we are in username check")	
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
        return &models.Account{Status:0,Message:"enterprise_username already exists"}, err

	} 


query = fmt.Sprintf("update enterprise SET enterprise_username='%s' , enterprise_pin='%s' where enterprise_id='%s' ", enterprise_username,enterprise_pin,enterprise_id)
_, err = c.DbConn.ExecContext(ctx, query)
fmt.Println("rows:%s",rows)
fmt.Println("err:%s",err)
if err != nil {
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching account_details from database")
//		return result[0], nil
	return &models.Account{Status:0,Message:"Failed"}, err
}
query = fmt.Sprintf("update account_login SET username='%s', password='%s' where enterprise_id='%s' ", enterprise_username,enterprise_pin,enterprise_id)
_, err = c.DbConn.ExecContext(ctx, query)
fmt.Println("rows:%s",rows)
fmt.Println("err:%s",err)
if err != nil {
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching account_details from database")
//		return result[0], nil
	return &models.Account{Status:0,Message:"Failed"}, err
}
fmt.Println("I am about to go in defer")
//add defer function
defer func() {
	fmt.Println("I am in defer")
err := rows.Close()
if err != nil {
	panic(err)
}
}()




/**result = make([]*models.Account, 0)
for rows.Next() {
     fmt.Println("I am in results")
	t := new(models.Account)

	t.Status=1
	t.Message="changes successfully registered"
	fmt.Println("hi:",t.Status)
	fmt.Println("hi:",t.Message)

	err := rows.Scan(&t.Message)
	fmt.Println("err while scan:%s",err)
	fmt.Println(t.Message)
	if err != nil {
		fmt.Println("here rows.scan is return errors")
		logger.Logger.WithError(err).WithField("query", query).
		Errorf("Error while fetching carriers from database")
//			return nil, err
		//return &models.Account{Status:0,Message:"email_id already exists"}, err
	}
	result = append(result, t)		
} **/


/**if len(result)== 0 {
	fmt.Println("here length  is  zero ")	
	logger.Logger.WithError(err).WithField("query", query).
	Errorf("Error while fetching carriers from database")
//		return &models.Account{Message:"success"}, nil
	return &models.Account{Status:0,Message:"Failed"}, err

} **/
  //return result[0],nil
  return &models.Account{Status:1,Message:"changes registered successfully"}, nil	
}

// show one Enterprise with join

func  (c *accountRepository) ShowOne_Enterprise(ctx context.Context,enterprise_id string) (*models.ShowJoin_Enterprise, error) {

	//	hashMd5 := md5.New()
	//	io.WriteString(hashMd5, secret)
	//	md5Hash := fmt.Sprintf("%x", hashMd5.Sum(nil))
	
	//fmt.Println("hello username repository", enterprise_username)
	query := fmt.Sprintf("SELECT account_login.is_superuser,account_login.is_active,enterprise.enterprise_email,enterprise.enterprise_username,enterprise.enterprise_pin,enterprise.enterprise_mobile,enterprise.created_at,enterprise.updated_at FROM account_login INNER JOIN enterprise ON account_login.enterprise_id = enterprise.enterprise_id WHERE enterprise.enterprise_id='%s'",enterprise_id)//md5Hash)
			 fmt.Printf("hi:%s", query)
			 rows, err := c.DbConn.QueryContext(ctx, query)
			 if err != nil {
			   logger.Logger.WithError(err).WithField("query", query).
				  Errorf("Error while fetching project_details By ID from database")
			      fmt.Println("repository/account.go error\n", err)
			      fmt.Println("repository/account.go rows\n", rows)
			      return &models.ShowJoin_Enterprise{Status:"0",Message:"Failed"}, nil
			  // return nil, err
			 }
	  
			if rows == nil {
			   fmt.Println("rows are nil\n", err)
			}
	  
			logger.Logger.Info("In accountRepo:GetProjectById")
			defer func() {
			   err := rows.Close()
			   if err != nil {
				  panic(err)
			   }
			}()
	  
			fmt.Printf("repository/account.go\n")
	  
			result := make([]*models.ShowJoin_Enterprise, 0)
			for rows.Next() {
			   t := new(models.ShowJoin_Enterprise)
			   //    err := rows.Scan(&t.ProjectId, &t.ProjectName, &t.ProjectDetails, &t.ProjectComment)
			  // t.Status=1
			   err := rows.Scan(&t.Is_superuser,&t.Is_active,&t.Email,&t.Username, &t.Pin, &t.Mobile,&t.Created ,&t.Updated)
	  
			   if err != nil {
				  logger.Logger.WithError(err).WithField("query", query).
					 Errorf("Error while fetching carriers from database")
				 // fmt.Printf("rows.next giving error\n", err)
				  return &models.ShowJoin_Enterprise{Status:"0",Message:"Failed"}, err
			   }
			   result = append(result, t)
			}
	  
			fmt.Println("result ", result)
	  
			if len(result) == 0 {
			   logger.Logger.WithError(err).WithField("query", query).
				  Errorf("Error while fetching carriers from database")
			   fmt.Println("len(result) is zero\n", err)
	  
			  // status := &models.{Status: "0", Msg: "Failed", ResponseCode: "200"}
			   //fmt.Printf("status \n", status)
			   //result = append(result, status)
			  // return result, err
			   return &models.ShowJoin_Enterprise{Status:"0",Message:"Failed"}, err
			   //    return &models.ProjectById{Status:"0",Msg:"Failed",ResponseCode:"200"}, nil
	  
			}
		
			return result[0], err
	}
 



					

					

			

					