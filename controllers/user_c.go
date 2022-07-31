package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	"encoding/json"
	_"strings"
	"log"
	_"time"
)
type DatatableAjaxUser struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.User `json:"data"`
}
type UserDatatable struct {
	Search string `json:"search"`
	Start int `json:"start"`
	Length int `json:"length"`
	Draw int `json:"draw"`
	Shortby string `json:"shortby"`
	Type string `json:"type"`
	Idcompany string `json:"idcompany"`
}
func GetUserForDatatable(c *gin.Context)(DatatableAjaxUser){
	// data:=CheckAPIAuthentication(c)
	// idrole,_:=strconv.Atoi(os.Getenv("IDROLE_SU")/)
	var user models.User
	jsonData,_ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	idcompany,_:=strconv.Atoi(c.Query("idcompany"))
	idrole,_:=strconv.Atoi(c.Query("idrole"))
	search:=c.Query("search[value]")
	start,_:=strconv.Atoi(c.Query("start"))
	length,_:=strconv.Atoi(c.Query("length"))
	draw,_:=strconv.Atoi(c.Query("draw"))
	colShort,_:=strconv.Atoi(c.Query("order[0][column]"))
	typ:=c.Query("order[0][dir]")
	shortby:=""
	switch colShort {
    case 0:
        shortby="id"
    case 1:
        shortby="username"
    case 2:
        shortby="email"
    case 3:
        shortby="idrole"
    case 4:
        shortby="idcompany"
    }
    var users [] models.User
    var total int64
    u:=ClaimToken(c)
    users,total=models.GetUsersForDatatable(search,start,length,shortby,typ,models.User{Idcompany:idcompany,Idapplication:u.Idapplication,Idrole:idrole})
    
	ajax:=DatatableAjaxUser{Draw:draw,TotalFiltered:len(users),Total:total,Data:users}
	return ajax
}
func GenerateUserOwnerExist(c *gin.Context)(newUser models.UserMini){
	// var company models.Company
	// jsonData, _ := c.GetRawData()
	// json.Unmarshal(jsonData,&company)
	// //Check this company have Owner?
	// var newRole models.Role
	// roleexist:=models.GetRole(models.Role{Idcompany:company.Id,Role:"Owner"})
	// if len(roleexist)==0{
	// 	//Jika role owner belum tersimpan
	// 	newRole=models.CreateRole(models.Role{Role:"Owner",Idcompany:company.Id,Desc:"Owner of "+company.Name+" Company"})
	// }else{
	// 	newRole=roleexist[0]
	// }
	// log.Println("This new role : ",newRole)
	// //Buat role jadi FullAccess
	// menu1:=models.GetMenu1CantAccess(newRole)
	// menu2:=models.GetMenu2CantAccess(newRole)
	// for i := 0; i < len(menu1); i++ {
	//     models.CreateAccess1(models.Access1{Create:true,Read:true,Update:true,Delete:true,Idmenu1:menu1[i].Id,Idrole:newRole.Id})
	    
	// }
	// log.Println("Full access menu 1 already given")
	// for i := 0; i < len(menu2); i++ {
	//     models.CreateAccess2(models.Access2{Create:true,Read:true,Update:true,Delete:true,Idmenu2:menu2[i].Id,Idrole:newRole.Id})

	// }
	// log.Println("Full access menu 2 already given")
	// app:=ClaimToken(c)
	// now:=time.Now().Local().Unix()
	// newPass:=strconv.FormatInt(now+123, 10)
	// log.Println("Ini passsss",now,newPass)
	// encrypted,_:=EncodingBcrypt(newPass)
	// // log.Println(models.GetUser(models.User{Idcompany:company.Id,Idrole:newRole.Id}))
	// uExist:=models.GetUser(models.User{Idcompany:company.Id,Idrole:newRole.Id,Idapplication:app.Id})
	// var newUser models.UserMini
	// if uExist.Username==""{
	// 	newUser=models.CreateUser(models.UserMini{Idcompany:company.Id,Idrole:newRole.Id,Username:"u-"+strconv.FormatInt(now,10),Password:encrypted,Email:company.Email,Mobilenumber:company.Phone})
	// }else{
	// 	newUser=models.UserMini{Id:uExist.Id,Idcompany:uExist.Idcompany,Idrole:uExist.Idrole,Username:uExist.Username,Password:uExist.Password,Email:uExist.Email,Mobilenumber:uExist.Mobilenumber}
	// }
	// uOld:=strings.Replace(newUser.Username, "u-", "", -1)
	// userx,_:= strconv.Atoi(uOld)
	// userx=userx+123
	// log.Println(uOld,userx)
	// newUser.Password=strconv.FormatInt(int64(userx),10)
	return newUser
}
func GetUser(c *gin.Context)(u []models.User){
	var user models.User
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)	
	user.Idapplication=app.Id
	usr:=models.GetUser(user)
	for i := 0; i < len(usr); i++ {
		usr[i].Password="***Mau ngintip password ya ???***"
	}
	return usr
}
func CreateUser(c *gin.Context)(u models.UserMini){
	var user models.UserMini
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	encoded:=""
	if user.Password == ""{
		encoded,_=EncodingBcrypt(user.Username)
		user.Password=encoded

	}else{
		encoded,_=EncodingBcrypt(user.Password)
		user.Password=encoded
	}	

	exist:=models.GetUser(models.User{Username:user.Username})
	if user.Username=="" || user.Idrole ==0 || user.Idcompany ==0 || user.Idapplication ==0{
		c.JSON(400,gin.H{"result":"username, idrole, idcompany, idapplication is mandatory"})
	} else if len(exist)!=0{
		c.JSON(400,gin.H{"result":"Username already taken"})
		
	}else{
		u=models.CreateUser(user)
	}
	
	u.Password="***Mau ngintip password ya ???***"
	return u
}

func UpdateUser(c *gin.Context)(u models.User){
	var user models.User
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	exist:=models.GetUser(models.User{Username:user.Username})
	if user.Id==0 || user.Username=="" || user.Idrole ==0 || user.Idcompany ==0 || user.Idapplication ==0{
		c.JSON(400,gin.H{"result":"id, username, idrole, idcompany, idapplication is mandatory"})
	} else if len(exist)!=0{
		c.JSON(400,gin.H{"result":"Username already taken"})
		
	}else{
		u=models.UpdateUser(user)
	}
	u.Password="***Mau ngintip password ya ???***"
	return u
}
func DeleteUser(c *gin.Context)(models.User){
	var user models.User	
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	user.Active=false
	log.Println(user,"errrrrrrrrrrrrrr")
	u:=models.DeleteUser(user)
	u.Password="***Mau ngintip password ya ???***"
	return u
}