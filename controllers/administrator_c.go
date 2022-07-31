package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	_"log"
	"errors"
	"encoding/json"
)
type DatatableAjaxPrefix struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Prefix `json:"data"`
}
type DatatableAjaxPricePlan struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.PricePlan `json:"data"`
}
type DatatableAjaxPricePlanDetail struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.PricePlanDetail `json:"data"`
}
func CheckAliasExist(idoperator int,p []string)(bool,error){
	exist:=false
	var err error
	aliasCount:=len(models.CheckAliasExist(idoperator,p))
	if(aliasCount>0){
		exist=true
		err=errors.New("Alias / Kode is already exist. "+string(aliasCount) +" found")
	}
	return exist,err
}
func CheckNamaExist(idoperator int,p []string)(bool,error){
	exist:=false
	var err error
	namaCount:=len(models.CheckNamaExist(idoperator,p))
	if(namaCount>0){
		exist=true
		err=errors.New("Nama is already exist. "+string(namaCount) +" found")
	}
	return exist,err
}

func GetPrefixForDatatable(c *gin.Context)(DatatableAjaxPrefix){
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
        shortby="alias"
    }
    var pref [] models.Prefix
    var total int64
    pref,total=models.GetPrefixForDatatable(search,start,length,shortby,typ,models.Prefix{})

	ajax:=DatatableAjaxPrefix{Draw:draw,TotalFiltered:len(pref),Total:total,Data:pref}
	return ajax
}

func checkExistPrefix(p models.Prefix)(bool,error){
	exist:=false
	var err error
	if(len(models.GetPrefix(models.Prefix{Name:p.Name}))>0){
		exist=true
		err=errors.New("Name is already exist")
	}else if len(models.GetPrefix(models.Prefix{Alias:p.Alias}))>0 {
		exist=true
		err=errors.New("Alias is already exist")
	}
	return exist,err
}

func CreatePrefix(c *gin.Context)(models.Prefix,error){
	var prefix models.Prefix
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prefix)
	exist,err:=checkExistPrefix(prefix)
	if exist{
		return models.Prefix{},err
	}else{
		return models.CreatePrefix(prefix),nil
	}
}

func GetPrefix(c *gin.Context)([]models.Prefix){
	var prefix models.Prefix
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prefix)
	return models.GetPrefix(prefix)
}
func UpdatePrefix(c *gin.Context)(models.Prefix){
	var prefix models.Prefix
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prefix)
	return models.UpdatePrefix(prefix)	
}

func DeletePrefix(c *gin.Context)(models.Prefix){
	var prefix [] models.Prefix
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prefix)
	return models.DeletePrefix(prefix)	
}