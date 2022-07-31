package models
import (
  "gorm.io/gorm"
  "hris/config"
)
var db *gorm.DB
func init(){
	config.Connect()
	db=config.GetDB()
}
type User struct {
	Idcompany int `json:"idcompany"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Role string `json:"role"`
	Company string `json:"company"`
	Idrole int `json:"idrole"`
	Idapplication int `json:"idapplication"`
	Id int `gorm:"primaryKey" json:"id"`
	Active bool `json:"active"`
}
type UserMini struct {
	Id int `gorm:"primaryKey" json:"id"`
	Idcompany int `json:"idcompany"`
	Idrole int `json:"idrole"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}
type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Idrole int `json:"idrole"`
	Idcompany int `json:"idcompany"`
}

func CreateUser(u UserMini) (UserMini){
	db.Table("tbl_user").Create(&u)
	return u
}
func UpdateUser(u User) (User){
	var usr User
	db.Table("tbl_user").Where("id=?",u.Id).Updates(u).Scan(&usr)
	return usr
}
func DeleteUser(u User) (res User){
	db.Table("tbl_user").Where("id=?",u.Id).Updates(map[string]interface{}{"active":u.Active}).Scan(&res)
	return res
}
func GetUser(usr User) (u []User){
	db.Table("tbl_user").Where(usr).Find(&u)
	return u
}
func GetSingle(usr User) (u User){ //Where
	db.Table("tbl_user").Where(usr).Find(&u)
	return u
}
func GetUsersForDatatable(search string,offset int, limit int, shortby string, typ string,usr User) (u []User,totalall int64){
	if search !=""{
		db.Table("tbl_user").Where(usr).Limit(limit).Offset(offset).Where("username like '%"+search+"%' or email like '%"+search+"%' or mobilenumber like '%"+search+"%'").Order(shortby+" "+typ).Scan(&u)
			
	}else{
		db.Table("tbl_user").Where(usr).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&u)
	}
	var users []User
	res :=db.Table("tbl_user").Where(usr).Scan(&users)
	return u,res.RowsAffected
}
func GetUserFull(user User) (User){
	var u User
	if user.Username != ""{
		db.Raw("select u.id,username,password,email,mobilenumber,idrole,idcompany,idapplication, (select role from tbl_role r where r.id=idrole ) as role,(select name from tbl_company c where c.id=idcompany) as company from tbl_user u where username=?",user.Username).Scan(&u)
	}else{
		db.Raw("select u.id,username,password,email,mobilenumber,idrole,idcompany,idapplication, (select role from tbl_role r where r.id=idrole ) as role,(select name from tbl_company c where c.id=idcompany) as company from tbl_user u where id=?",user.Id).Scan(&u)
	}
	
	return u
}