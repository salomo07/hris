package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	"log"
	"encoding/json"
)
type DatatableAjaxRole struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Role `json:"data"`
}
func GetRolesbyCompany(idcompany int)([]models.Role){
	return models.GetRolesbyCompany(idcompany)
}
func GetRole(c *gin.Context)([]models.Role){
	var role models.Role
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&role)
	role.Idapplication=app.Id
	r:=models.GetRole(role)
	return r
}
func CreateRole(c *gin.Context)(r models.Role){
	var role models.Role
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&role)
	role.Idapplication=app.Id
	role.Active=true
	if role.Role!="" && role.Idcompany!=0 {
		r=models.CreateRole(role)
	}else{
		c.JSON(400, gin.H{"result":"role, idcompany is mandatory"})
	}
	return r
}

func UpdateRole(c *gin.Context)(r models.Role){
	var role models.Role
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&role)
	role.Idapplication=app.Id
	if role.Id!=0 && role.Role!="" && role.Idcompany!=0 {
		r=models.UpdateRole(role)
	}else{
		c.JSON(400, gin.H{"result":"id,role, idcompany is mandatory"})
	}
	return r
}
func DeleteRole(c *gin.Context)(r models.Role){
	var role models.Role
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&role)
	role.Active=false
	log.Println(role)
	if role.Id!=0 {
		r=models.DeleteRole(role)
	}else{
		c.JSON(400, gin.H{"result":"id is mandatory"})
	}
	return r
}

func GetRoleForDatatable(c *gin.Context)(DatatableAjaxRole){
	app:=CheckAPI(c)
	idcompany,_:=strconv.Atoi(c.Query("idcompany"))
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
        shortby="idcompany"
    case 2:
        shortby="role"
    case 3:
        shortby="desc"
    }
    var roles [] models.Role
    var total int64
	roles,total=models.GetRoleForDatatable(search,start,length,shortby,typ,models.Role{Idcompany:idcompany,Idapplication:app.Id})
    
	ajax:=DatatableAjaxRole{Draw:draw,TotalFiltered:len(roles),Total:total,Data:roles}
	return ajax
}