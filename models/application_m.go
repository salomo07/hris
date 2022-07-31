package models
import (
  _"log"
)
type Application struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Startdate string `json:"startdate"`
	Enddate string `json:"enddate"`
	Typekey string `json:"typekey"`
	Key string `json:"key"`
}
func CreateApplication(a Application) (Application){
	db.Table("tbl_application").Create(&a)
	return a
}
func GetApplication(a Application)(r []Application){
	db.Table("tbl_application").Where(a).Find(&r)
	return r
}
func UpdateApplication(a Application) (res Application){
	db.Table("tbl_application").Where("id=?",a.Id).Updates(Application{Name:a.Name,Phone:a.Phone,Address:a.Address,Startdate:a.Startdate,Enddate:a.Enddate,Typekey:a.Typekey,Key:a.Key}).Scan(&res)
	return res
}
func DeleteApplication(a []Application) (Application){
	var res Application
	for i := 0; i < len(a); i++ {
	    db.Table("tbl_application").Where(a[i]).Delete(&res)
	}
	return res
}
// func CreateApiKey(api ApiKey) (ApiKey){
// 	var a ApiKey
// 	db.Table("tbl_apikey").Where("name = ? and type=?",api.Name,api.Type).Find(&a)
// 	if a.Name !=""{
// 		db.Table("tbl_apikey").Where("id = ?",a.Id).Updates(api).Find(&api)
// 	}else{
// 		db.Table("tbl_apikey").Create(&api)
// 	}
// 	return api
// }
// func CreateBearerKey(api ApiKey) (ApiKey){
// 	var a ApiKey
// 	db.Table("tbl_apikey").Where("name = ? and type=?",api.Name,api.Type).Find(&a)
// 	if a.Name !=""{
// 		db.Table("tbl_apikey").Where("id = ?",a.Id).Updates(api).Find(&api)
// 	}else{
// 		db.Table("tbl_apikey").Create(&api)
// 	}

// 	return api
// }
func GetAPIKey(app Application) (a Application){
	db.Table("tbl_application").Where(app).Find(&a)
	return a
}
