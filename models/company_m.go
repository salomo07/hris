package models
import (
  _"log"
)
type Company struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Level string `json:"level"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}

func GetCompany(c Company)(r []Company){
	db.Table("tbl_company").Where(c).Where(Company{Active:c.Active}).Find(&r)
	return r
}
func GetCompanyForDatatable(search string,offset int, limit int, shortby string, typ string, idapplication int) (c []Company,totalall int64){
	if search !=""{
		db.Table("tbl_company").Limit(limit).Offset(offset).Where("name like '%"+search+"%' or phone like '%"+search+"%'").Where(Company{Idapplication:idapplication}).Order(shortby+" "+typ).Scan(&c)
			
	}else{
		db.Table("tbl_company").Where(Company{Idapplication:idapplication}).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&c)
	}
	var companies []Company
	res :=db.Table("tbl_company").Where(Company{Idapplication:idapplication}).Scan(&companies)
	return c,res.RowsAffected
}
func GetApplicationForDatatable(search string,offset int, limit int, shortby string, typ string) (a []Application,totalall int64){
	if search !=""{
		db.Table("tbl_application").Limit(limit).Offset(offset).Where("name like '%"+search+"%' or phone like '%"+search+"%' or address like '%"+search+"%'").Order(shortby+" "+typ).Scan(&a)
			
	}else{
		db.Table("tbl_application").Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&a)
	}
	var app []Application
	res :=db.Table("tbl_application").Scan(&app)
	return a,res.RowsAffected
}
func CreateCompany(c Company) (Company){
	c.Active=true
	db.Table("tbl_company").Create(&c)
	return c
}
func UpdateCompany(c Company) (res Company){
	db.Table("tbl_company").Where("id=?",c.Id).Updates(map[string]interface{}{"id":c.Id,"name":c.Name,"phone":c.Phone,"level":c.Level,"email":c.Email,"idapplication":c.Idapplication,"active":c.Active}).Scan(&res)
	return res
}
func DeleteCompany(c Company) (res Company){
	db.Table("tbl_company").Where("id=?",c.Id).Updates(map[string]interface{}{"active":false}).Scan(&res)
	return res
}