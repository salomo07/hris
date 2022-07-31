package models
import (
  "log"
)
type PricePlan struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
	Mintiketdeposit int `json:"mintiketdeposit"`
	Maxselisih int `json:"maxselisih"`
	Sistempoin bool `json:"sistempoin"`
	Idcompany int `json:"idcompany"`
}
type PricePlanDetail struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idpriceplan int `json:"idpriceplan"`
	Idprefix int `json:"idprefix"`
	Idproduk int `json:"idproduk"`
	Nama string `json:"nama"`
	Hargajual int `json:"hargajual"`
	Komisi int `json:"komisi"`
	Margin int `json:"margin"`
}
type PricePlanDetailInsert struct {
	Idpriceplan int `json:"idpriceplan"`
	Idproduk int `json:"idproduk"`
	Hargajual int `json:"hargajual"`
	Komisi int `json:"komisi"`
	Margin int `json:"margin"`
}
func GetPricePlanDetailForDatatable(idpriceplan int) (rslt []PricePlanDetail,totalall int64){
	db.Table("vpriceplandetail").Where("idpriceplan=?",idpriceplan).Scan(&rslt)	
	var det []PricePlanDetail
	res :=db.Table("vpriceplandetail").Scan(&det)
	log.Println(rslt)

	return rslt,res.RowsAffected
}
func UpsertPricePlanDetail(p PricePlanDetail)(string){
	var exist PricePlanDetail
	db.Table("tbl_priceplandetail").Where(PricePlanDetail{Idpriceplan:p.Idpriceplan,Idproduk:p.Idproduk}).Scan(&exist)
	log.Println(exist)
	if exist.Id !=0{
		db.Table("tbl_priceplandetail").Where("id=?",exist.Id).Updates(map[string]interface{}{"hargajual":p.Hargajual,"komisi":p.Komisi,"margin":p.Margin})
		log.Println("Updated : ",p)
	}else{
		log.Println("Insert : ",p.Idpriceplan,p.Idproduk,p.Hargajual)
		data:=PricePlanDetailInsert{Idpriceplan:p.Idpriceplan,Idproduk:p.Idproduk,Hargajual:p.Hargajual,Komisi:p.Komisi,Margin:p.Margin}
		db.Table("tbl_priceplandetail").Create(data)
	}
	return ""
}
func GetPricePlanForDatatable(search string,offset int, limit int, shortby string, typ string,price PricePlan) (rslt []PricePlan,totalall int64){
	db.Table("tbl_priceplan").Where("idcompany =?",price.Idcompany).Where(price).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	var pp []PricePlan
	res :=db.Table("tbl_priceplan").Where("idcompany =?",price.Idcompany).Scan(&pp)
	return rslt,res.RowsAffected
}
func CreatePricePlan(p PricePlan)(PricePlan){
	db.Table("tbl_priceplan").Create(&p)
	return p
}
func UpdatePricePlan(p PricePlan)(res PricePlan){
	db.Table("tbl_priceplan").Where("id=?",p.Id).Updates(map[string]interface{}{"nama":p.Nama,"mintiketdeposit":p.Mintiketdeposit,"maxselisih":p.Maxselisih,"sistempoin":p.Sistempoin}).Scan(&res)
	return res
}
func GetPricePlan(p PricePlan)(res []PricePlan){
	db.Table("tbl_priceplan").Where(p).Find(&res)
	return res
}
func DeletePricePlan(p []PricePlan)(PricePlan){
	var res PricePlan
	for i := 0; i < len(p); i++ {
	    db.Table("tbl_priceplan").Where(p[i]).Delete(&res)
	}
	return res
}