package models
import (
  "log"
)
type Access1 struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Create bool `json:"create"`
	Read bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
	Idmenu1 int `json:"idmenu1"`
	Idrole int `json:"idrole"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}

type Access2 struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Create bool `json:"create"`
	Read bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
	Idmenu2 int `json:"idmenu2"`
	Idrole int `json:"idrole"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}

func GetAccess2(a Access2)(res []Access2){
	if(a.Idrole!=0){
		db.Table("tbl_access2").Where(a).Scan(&res)
	}
	return res
}
func GetAccess1(a Access1)(res []Access1){
	log.Println(a)
	if(a.Idrole!=0){
		db.Table("tbl_access1").Where(a).Scan(&res)
	}
	return res
}
func GetAccess1ForDatatable(search string,offset int, limit int, acc Access1) (rslt []Access1,totalall int64){
	if search !=""{
		db.Table("tbl_access1").Where(acc).Limit(limit).Offset(offset).Order("idmenu1 asc").Scan(&rslt)
			
	}else{
		db.Table("tbl_access1x").Where(acc).Limit(limit).Offset(offset).Order("idmenu1 asc").Scan(&rslt)
	}
	var accesss []Access1
	res :=db.Table("tbl_access1").Scan(&accesss)
	return rslt,res.RowsAffected
}
func GetAccess2ForDatatable(search string,offset int, limit int,acc Access2) (rslt []Access2,totalall int64){
	if search !=""{
		db.Table("tbl_access2").Where(acc).Limit(limit).Offset(offset).Order("idmenu2 asc").Scan(&rslt)
			
	}else{
		db.Table("tbl_access2").Where(acc).Limit(limit).Offset(offset).Order("idmenu2 asc").Scan(&rslt)
	}
	var menus []Access2
	res :=db.Table("tbl_access2").Scan(&menus)
	return rslt,res.RowsAffected
}
func CreateAccess1(a Access1)(Access1){
	db.Table("tbl_access1").Create(&a)
	return a
}
func CreateAccess2(a Access2)(Access2){
	db.Table("tbl_access2").Create(&a)
	return a
}
func UpdateAccess1(a Access1) (res Access1){
	db.Table("tbl_access1").Where("id=?",a.Id).Updates(map[string]interface{}{"read":a.Read,"create":a.Create,"update":a.Update,"delete":a.Delete}).Scan(&res)
	return res
}
func UpdateAccess2(a Access2) (res Access2){
	db.Table("tbl_access2").Where("id=?",a.Id).Updates(map[string]interface{}{"read":a.Read,"create":a.Create,"update":a.Update,"delete":a.Delete}).Scan(&res)
	return res
}
// func DeleteAccess1(a Access1) (res Access1){
// 	db.Table("tbl_access1").Where("id=?",a.Id).Updates(map[string]interface{}{"active":a.Active}).Scan(&res)
// 	return res
// }
// func DeleteAccess2(a Access2) (res Access2){
// 	db.Table("tbl_access2").Where("id=?",a.Id).Updates(map[string]interface{}{"active":a.Active}).Scan(&res)
// 	return res
// }