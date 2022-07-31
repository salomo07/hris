package models
import (
  _"log"
)
type JenisProduk struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idcompany int `json:"idcompany"`
	Nama string `json:"nama"`
}

type StatusProduk struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idcompany int `json:"idcompany"`
	Nama string `json:"nama"`
}
type Produk struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idprefix int `json:"idprefix"`
	Nama string `json:"nama"`
	Nominal int `json:"nominal"`
	Alias string `json:"alias"`
	Idjenis int `json:"idjenis"`
	Idstatus int `json:"idstatus"`
	Infogangguan string `json:"infogangguan"`
	Deskripsi string `json:"deskripsi"`
}

func GetProdukForDatatable(search string,offset int, limit int, shortby string, typ string,pref Produk) (rslt []Produk,totalall int64){
	if search !=""{
		db.Table("tbl_produk").Where(pref).Limit(limit).Offset(offset).Where("nama like '%"+search+"%' or alias like '%"+search+"%' or cast( nominal as varchar) like '%"+search+"%'").Order("idprefix,id asc").Scan(&rslt)
			
	}else{
		db.Table("tbl_produk").Where(pref).Limit(limit).Offset(offset).Order("idprefix,id asc").Scan(&rslt)
	}
	var jenis []Produk
	res :=db.Table("tbl_produk").Scan(&jenis)
	return rslt,res.RowsAffected
}
func GetStatusProdukForDatatable(search string,offset int, limit int, shortby string, typ string,pref StatusProduk) (rslt []StatusProduk,totalall int64){
	if search !=""{
		db.Table("tbl_ref_statusproduk").Where(pref).Limit(limit).Offset(offset).Where("nama like '%"+search+"%'").Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_ref_statusproduk").Where(pref).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var status []StatusProduk
	res :=db.Table("tbl_ref_statusproduk").Scan(&status)
	return rslt,res.RowsAffected
}
func GetJenisProdukForDatatable(search string,offset int, limit int, shortby string, typ string,pref JenisProduk) (rslt []JenisProduk,totalall int64){
	if search !=""{
		db.Table("tbl_ref_jenisproduk").Where(pref).Limit(limit).Offset(offset).Where("nama like '%"+search+"%'").Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_ref_jenisproduk").Where(pref).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var jenis []JenisProduk
	res :=db.Table("tbl_ref_jenisproduk").Scan(&jenis)
	return rslt,res.RowsAffected
}

func CreateJenisProduk(p JenisProduk)(JenisProduk){
	db.Table("tbl_ref_jenisproduk").Create(&p)
	return p
}
func CreateStatusProduk(p StatusProduk)(StatusProduk){
	db.Table("tbl_ref_statusproduk").Create(&p)
	return p
}
func CreateProduk(p Produk)(Produk){
	db.Table("tbl_produk").Create(&p)
	return p
}

func UpdateProduk(p Produk)(res Produk){
	db.Table("tbl_produk").Where("id=?",p.Id).Updates(map[string]interface{}{"nama":p.Nama,"alias":p.Alias,"idjenis":p.Idjenis,"infogangguan":p.Infogangguan,"deskripsi":p.Deskripsi,"idstatus":p.Idstatus,"nominal":p.Nominal,"idprefix":p.Idprefix}).Scan(&res)
	return res
}
func UpdateJenisProduk(p JenisProduk)(res JenisProduk){
	db.Table("tbl_ref_jenisproduk").Where("id=?",p.Id).Updates(map[string]interface{}{"nama":p.Nama}).Scan(&res)
	return res
}
func UpdateStatusProduk(p StatusProduk)(res StatusProduk){
	db.Table("tbl_ref_statusproduk").Where("id=?",p.Id).Updates(map[string]interface{}{"nama":p.Nama}).Scan(&res)
	return res
}
func CheckAliasExist(idoperator int,p []string)(res []Produk){
	db.Raw("select * from tbl_produk where idprefix=? and alias in (?)",idoperator,p).Scan(&res)
	return res
}
func CheckNamaExist(idoperator int,p []string)(res []Produk){
	db.Raw("select * from tbl_produk where idprefix=? and nama in (?)",idoperator,p).Scan(&res)
	return res
}
func GetProduk(p Produk)(res []Produk){
	db.Table("tbl_produk").Where(p).Find(&res)
	return res
}
func GetJenisProduk(p JenisProduk)(res []JenisProduk){
	db.Table("tbl_ref_jenisproduk").Where(p).Find(&res)
	return res
}
func GetStatusProduk(p StatusProduk)(res []StatusProduk){
	db.Table("tbl_ref_statusproduk").Where(p).Find(&res)
	return res
}
func DeleteJenisProduk(p []JenisProduk)(JenisProduk){
	var res JenisProduk
	for i := 0; i < len(p); i++ {
	    db.Table("tbl_ref_jenisproduk").Where(p[i]).Delete(&res)
	}
	return res
}
func DeleteStatusProduk(p []StatusProduk)(StatusProduk){
	var res StatusProduk
	for i := 0; i < len(p); i++ {
	    db.Table("tbl_ref_statusproduk").Where(p[i]).Delete(&res)
	}
	return res
}
func DeleteProduk(p []Produk)(Produk){
	var res Produk
	for i := 0; i < len(p); i++ {
	    db.Table("tbl_produk").Where(p[i]).Delete(&res)
	}
	return res
}