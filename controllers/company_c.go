package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	_"log"
	"encoding/json"
)
type DatatableAjaxCompany struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Company `json:"data"`
}
type DatatableAjaxApplication struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Application `json:"data"`
}

func GetCompanyForDatatable(c *gin.Context)(DatatableAjaxCompany){
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
        shortby="phone"
    case 3:
        shortby="level"
    }
    user:=ClaimToken(c)
    var companies [] models.Company
    var total int64
    companies,total=models.GetCompanyForDatatable(search,start,length,shortby,typ,user.Idapplication)
	ajax:=DatatableAjaxCompany{Draw:draw,TotalFiltered:len(companies),Total:total,Data:companies}
	return ajax
}
func GetCompany(c *gin.Context)(com []models.Company){
	var company models.Company
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&company)
	company.Idapplication=app.Id
	com=models.GetCompany(company)
	return com
}
func CreateCompany(c *gin.Context)(companyRes models.Company){
	var company models.Company
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&company)
	company.Idapplication=app.Id
	if company.Name!="" && company.Phone!=""{
		companyRes=models.CreateCompany(company)
	}else{
		c.JSON(400, gin.H{"result":"name, phone is mandatory"})
	}
	return companyRes
}
func UpdateCompany(c *gin.Context)(companyRes models.Company){
	var company models.Company	
	app:=CheckAPI(c)
	jsonData, _ := c.GetRawData()
	json.Unmarshal([]byte(jsonData),&company)
	company.Idapplication=app.Id
	if company.Id!=0 && company.Name!="" && company.Phone!=""{
		companyRes=models.UpdateCompany(company)
	}else{
		c.JSON(400, gin.H{"result":"id, name, phone is mandatory"})
	}
	return companyRes
}
func DeleteCompany(c *gin.Context)(companyRes models.Company){
	var company models.Company	
	jsonData, _ := c.GetRawData()
	json.Unmarshal([]byte(jsonData),&company)
	
	if company.Id!=0{
		companyRes=models.DeleteCompany(company)
	}else{
		c.JSON(400, gin.H{"result":"id is mandatory"})
	}
	return companyRes
}
func GetApplicationForDatatable(c *gin.Context)(DatatableAjaxApplication){
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
        shortby="phone"
    case 3:
        shortby="address"
    }
    var app [] models.Application
    var total int64
    app,total=models.GetApplicationForDatatable(search,start,length,shortby,typ)
	ajax:=DatatableAjaxApplication{Draw:draw,TotalFiltered:len(app),Total:total,Data:app}
	return ajax
}
func CreateApplication(c *gin.Context)(models.Application){
	var app models.Application
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&app)
	return models.CreateApplication(app)
}
func UpdateApplication(c *gin.Context)(models.Application){
	var app models.Application	
	jsonData, _ := c.GetRawData()
	json.Unmarshal([]byte(jsonData),&app)
	return models.UpdateApplication(app)
}
