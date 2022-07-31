package models
import (
  _"log"
)
type Role struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idcompany int `json:"idcompany"`
	Role string `json:"role"`
	Desc string `json:"desc"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}
func GetRolesbyCompany(idcompany int)(r []Role){
	db.Table("tbl_role").Where(Role{Idcompany:idcompany}).Find(&r)
	return r
}
func GetRole(r Role)(res []Role){
	db.Table("tbl_role").Where(r).Find(&res)
	return res
}
func GetRoleForDatatable(search string,offset int, limit int, shortby string, typ string,rol Role) (rslt []Role,totalall int64){
	if search !=""{
		db.Table("tbl_role").Where("role like '%"+search+"%' or tbl_role.desc like '%"+search+"%'").Where(rol).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_role").Where(rol).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var roles []Role
	res :=db.Table("tbl_role").Where(rol).Scan(&roles)
	return rslt,res.RowsAffected
}
func CreateRole(r Role)(res Role){
	var rr Role
	db.Table("tbl_role").Where("idcompany=? and role=?",r.Idcompany,r.Role).Scan(&rr)
	//Cek jika role yg akan masuk sdh ada ditabel, tidak akan tersimpan
	if rr.Id ==0{
		db.Table("tbl_role").Create(&r)
	}
	return r
}
func UpdateRole(r Role) (res Role){
	db.Table("tbl_role").Where("id=?",r.Id).Updates(map[string]interface{}{"idcompany":r.Idcompany,"role":r.Role,"desc":r.Desc,"active":r.Active}).Scan(&res)
	return res
}
func DeleteRole(r Role) (res Role){
	db.Table("tbl_role").Where("id=?",r.Id).Updates(map[string]interface{}{"active":r.Active}).Scan(&res)
	return res
}
// func DeleteRole(r []Role) (Role){
// 	var res Role
// 	for i := 0; i < len(r); i++ {
// 	    db.Table("tbl_role").Where(r[i]).Delete(&res)
// 	}
// 	return res
// }