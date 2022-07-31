package models
import (
  _"log"
)
type Menu struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
	Url string `json:"url"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}
type Submenu struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idmenu1 int `json:"idmenu1"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
	Url string `json:"url"`
	Idapplication int `json:"idapplication"`
	Active bool `json:"active"`
}
type MenuAccess1 struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
	Url string `json:"url"`
	Create bool `json:"create"`
	Read bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
type SubmenuAccess2 struct {
	Id int  `gorm:"primaryKey" json:"id"`
	Idmenu1 int `json:"idmenu1"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
	Url string `json:"url"`
	Create bool `json:"create"`
	Read bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
func GetSubmenu(s Submenu)([]Submenu){
	var rslt []Submenu
	if s!=(Submenu{}){
		db.Table("tbl_menu2").Where(s).Order("id asc").Scan(&rslt)
	}else{
		db.Table("tbl_menu2").Order("id asc").Scan(&rslt)
	}
	return rslt
}
func GetMenu(m Menu)([]Menu){
	var rslt []Menu
	if m!=(Menu{}){
		db.Table("tbl_menu1").Where(m).Order("id asc").Scan(&rslt)
	}else{
		db.Table("tbl_menu1").Order("id asc").Scan(&rslt)
	}	
	return rslt
}
func GetMenuForLogin(idrole int,idapplication int)(res []MenuAccess1){
	db.Raw("SELECT m.id,name,icon,m.desc,m.url,a.create,a.read,a.update,a.delete FROM public.tbl_access1 a join tbl_menu1 m on m.id=a.idmenu1 where idrole=? and m.idapplication=? and a.read=true order by id asc",idrole,idapplication).Scan(&res)
	return res
}
func GetSubmenuForLogin(idrole int,idapplication int)(res []SubmenuAccess2){
	db.Raw("SELECT m.id,name,icon,m.desc,m.url,m.idmenu1,a.read,a.create,a.update,a.delete FROM public.tbl_access2 a join tbl_menu2 m on m.id=a.idmenu2 where idrole=? and m.idapplication=? and a.read=true order by m.id,m.idmenu1 asc",idrole,idapplication).Scan(&res)
	return res
}
func GetMenuForDatatable(search string,offset int, limit int, shortby string, typ string,menu Menu) (rslt []Menu,totalall int64){
	if search !=""{
		db.Table("tbl_menu1").Where(menu).Limit(limit).Offset(offset).Where("name like '%"+search+"%' or tbl_menu1.desc like '%"+search+"%' or icon like '%"+search+"%' or url like '%"+search+"%'").Order(shortby+" "+typ).Scan(&rslt)
			
	}else{
		db.Table("tbl_menu1").Where(menu).Limit(limit).Offset(offset).Order(shortby+" "+typ).Scan(&rslt)
	}
	var menus []Menu
	res :=db.Table("tbl_menu1").Scan(&menus)
	return rslt,res.RowsAffected
}
func GetSubmenuForDatatable(search string,offset int, limit int,menu Submenu) (rslt []Submenu,totalall int64){
	if search !=""{
		db.Table("tbl_menu2").Where(menu).Limit(limit).Offset(offset).Where("name like '%"+search+"%' or tbl_menu2.desc like '%"+search+"%' or icon like '%"+search+"%' or url like '%"+search+"%'").Order("idmenu1 asc").Scan(&rslt)
			
	}else{
		db.Table("tbl_menu2").Where(menu).Limit(limit).Offset(offset).Order("idmenu1 asc").Scan(&rslt)
	}
	var menus []Submenu
	res :=db.Table("tbl_menu2").Scan(&menus)
	return rslt,res.RowsAffected
}
func CreateMenu(r Menu)(Menu){
	db.Table("tbl_menu1").Create(&r)
	return r
}
func CreateSubmenu(r Submenu)(Submenu){
	db.Table("tbl_menu2").Create(&r)
	return r
}
func UpdateMenu(m Menu) (res Menu){
	db.Table("tbl_menu1").Where("id=?",m.Id).Updates(map[string]interface{}{"name":m.Name,"icon":m.Icon,"desc":m.Desc,"url":m.Url,"active":m.Active}).Scan(&res)
	return res
}
func UpdateSubmenu(m Submenu) (res Submenu){
	db.Table("tbl_menu2").Where("id=?",m.Id).Updates(map[string]interface{}{"name":m.Name,"icon":m.Icon,"desc":m.Desc,"url":m.Url,"idmenu1":m.Idmenu1,"active":m.Active}).Scan(&res)
	return res
}
// func DeleteMenu(r []Menu) (Menu){
// 	var res Menu
// 	for i := 0; i < len(r); i++ {
// 	    db.Table("tbl_menu1").Where(r[i]).Delete(&res)
// 	}
// 	return res
// }
// func DeleteSubmenu(r []Submenu) (Submenu){
// 	var res Submenu
// 	for i := 0; i < len(r); i++ {
// 	    db.Table("tbl_menu2").Where(r[i]).Delete(&res)
// 	}
// 	return res
// }
func GetMenu1CantAccess(a Role)(res []Menu){
	if(a.Id!=0){
		db.Raw("select * from tbl_menu1 where id not in (select idmenu1 from tbl_access1 where idrole=?)",a.Id).Scan(&res)
	}
	return res
}
func GetMenu2CantAccess(a Role)(res []Submenu){
	if(a.Id!=0){
		db.Raw("select * from tbl_menu2 where id not in (select idmenu2 from tbl_access2 where idrole=?)",a.Id).Scan(&res)
	}
	return res
}