package models
import (
  "log"
)
type Reseller struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
	Idupline int `json:"idupline"`
	Idpriceplan int `json:"idpriceplan"`
	Alamat string `json:"alamat"`
	Namapemilik string `json:"namapemilik"`
	Telppemilik string `json:"telppemilik"`
	Emailpemilik string `json:"emailpemilik"`
	Urlreport string `json:"urlreport"`
	Markupharga int `json:"markupharga"`
	Smspromo string `json:"smspromo"`
	Active bool `json:"active"`
}
type OwnerReseller struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idreseller int `json:"idreseller"`
	Iduser int `json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobilephone string `json:"mobilephone"`
	Idrole int `json:"idrole"`
	Idcompany int `json:"idcompany"`
}
func GetReseller(r Reseller)(res []Reseller){
	db.Table("tbl_reseller").Where(r).Find(&res)
	return res
}
func GetOwnerReseller(r OwnerReseller)(res []OwnerReseller){
	log.Println(r.Id)
	db.Raw("select r.id,idreseller,iduser,username,password,email,mobilenumber,idrole from tbl_userreseller r join tbl_user u on u.id=r.iduser join tbl_role rr on idrole=rr.id where idreseller=? and rr.role='Owner'",r.Id).Scan(&res)
	return res
}

func GetResellerForDatatable(search string,offset int, limit int, shortby string, typ string,rol Reseller) (rslt []Reseller,totalall int64){
	if search !=""{
		db.Table("tbl_reseller").Where(rol).Limit(limit).Offset(offset).Where("nama like '%"+search+"%' or alamat like '%"+search+"%' or namapemilik like '%"+search+"%'").Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_reseller").Where(rol).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var roles []Reseller
	res :=db.Table("tbl_reseller").Scan(&roles)
	return rslt,res.RowsAffected
}
func CreateReseller(r Reseller)(Reseller){
	db.Table("tbl_reseller").Create(&r)
	return r
}
func UpdateReseller(r Reseller) (res Reseller){
	db.Table("tbl_reseller").Where("id=?",r.Id).Updates(map[string]interface{}{"nama":r.Nama,"idupline":r.Idupline,"idpriceplan":r.Idpriceplan,"alamat":r.Alamat,"namapemilik":r.Namapemilik,"telppemilik":r.Telppemilik,"emailpemilik":r.Emailpemilik,"urlreport":r.Urlreport,"markupharga":r.Markupharga,"smspromo":r.Smspromo,"active":r.Active}).Scan(&res)
	return res
}
func DeleteReseller(r []Reseller) (Reseller){
	var res Reseller
	for i := 0; i < len(r); i++ {
	    db.Table("tbl_reseller").Where(r[i]).Delete(&res)
	}
	return res
}