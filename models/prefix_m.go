package models
import (
  _"log"
)
type Prefix struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	Prefixno string `json:"prefixno"`
	Mindigit int `json:"mindigit"`
	Maxdigit int `json:"maxdigit"`
	Publishweb bool `json:"publishweb"`
	Publishapp bool `json:"publishapp"`
	Imageurl string `json:"imageurl"`
}
func GetPrefixForDatatable(search string,offset int, limit int, shortby string, typ string,pref Prefix) (rslt []Prefix,totalall int64){
	if search !=""{
		db.Table("tbl_prefix").Where(pref).Limit(limit).Offset(offset).Where("name like '%"+search+"%' or alias like '%"+search+"%' or prefixno like '%"+search+"%'").Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_prefix").Where(pref).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var prefixs []Prefix
	res :=db.Table("tbl_prefix").Scan(&prefixs)
	return rslt,res.RowsAffected
}
func CreatePrefix(p Prefix)(Prefix){
	db.Table("tbl_prefix").Create(&p)
	return p
}
func UpdatePrefix(p Prefix)(res Prefix){

	db.Table("tbl_prefix").Where("id=?",p.Id).Updates(map[string]interface{}{"name":p.Name,"alias":p.Alias,"prefixno":p.Prefixno,"mindigit":p.Mindigit,"maxdigit":p.Maxdigit,"publishweb":p.Publishweb,"publishapp":p.Publishapp,"imageurl":p.Imageurl}).Scan(&res)
	return res
}
func GetPrefix(p Prefix)(res []Prefix){
	db.Table("tbl_prefix").Where(p).Find(&res)
	return res
}
func DeletePrefix(p []Prefix)(Prefix){
	var res Prefix
	for i := 0; i < len(p); i++ {
	    db.Table("tbl_prefix").Where(p[i]).Delete(&res)
	}
	return res
}