package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	"encoding/json"
)
type DatatableAjaxMenu struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Menu `json:"data"`
}
type DatatableAjaxSubmenu struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Submenu `json:"data"`
}
func CreateMenu(c *gin.Context)(models.Menu){
	var menu models.Menu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&menu)
	r:=models.CreateMenu(menu)
	return r
}
func CreateSubmenu(c *gin.Context)(models.Submenu){
	var submenu models.Submenu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&submenu)
	r:=models.CreateSubmenu(submenu)
	return r
}
func UpdateMenu(c *gin.Context)(models.Menu){
	var menu models.Menu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&menu)
	return models.UpdateMenu(menu)
}
func UpdateSubmenu(c *gin.Context)(models.Submenu){
	var submenu models.Submenu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&submenu)
	return models.UpdateSubmenu(submenu)
}
func DeleteMenu(c *gin.Context)(models.Menu){
	var menu models.Menu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&menu)
	menu.Active=false
	return models.UpdateMenu(menu)
}
func DeleteSubmenu(c *gin.Context)(models.Submenu){
	var submenu models.Submenu
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&submenu)
	submenu.Active=false
	return models.UpdateSubmenu(submenu)
}
func GetMenu ()([]models.Menu){
	menus:=models.GetMenu(models.Menu{})
	return menus
}
func GetSubmenu ()([]models.Submenu){
	sub:=models.GetSubmenu(models.Submenu{})
	return sub
}
func GetMenuForDatatable(c *gin.Context)(DatatableAjaxMenu){
	search:=c.Query("search[value]")
	start,_:=strconv.Atoi(c.Query("start"))
	length,_:=strconv.Atoi(c.Query("length"))
	draw,_:=strconv.Atoi(c.Query("draw"))
	colShort,_:=strconv.Atoi(c.Query("order[0][column]"))
	typ:=c.Query("order[0][dir]")
	shortby:=""
	switch colShort {
    case 0:
        shortby="id"
    case 1:
        shortby="name"
    case 2:
        shortby="url"
    case 3:
        shortby="icon"
    case 4:
        shortby="desc"
    }
    var menus [] models.Menu
    var total int64
    menus,total=models.GetMenuForDatatable(search,start,length,shortby,typ,models.Menu{})
	ajax:=DatatableAjaxMenu{Draw:draw,TotalFiltered:len(menus),Total:total,Data:menus}
	return ajax
}
func GetSubmenuForDatatable(c *gin.Context)(DatatableAjaxSubmenu){
	search:=c.Query("search[value]")
	start,_:=strconv.Atoi(c.Query("start"))
	length,_:=strconv.Atoi(c.Query("length"))
	draw,_:=strconv.Atoi(c.Query("draw"))
	
	
    var menus [] models.Submenu
    var total int64
    menus,total=models.GetSubmenuForDatatable(search,start,length,models.Submenu{})
	ajax:=DatatableAjaxSubmenu{Draw:draw,TotalFiltered:len(menus),Total:total,Data:menus}
	return ajax
}