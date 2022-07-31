package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	_"os"
	"strconv"
	_"log"
	"encoding/json"
)
type DatatableAjaxReseller struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Reseller `json:"data"`
}

func CreateReseller(c *gin.Context)(models.Reseller){
	var reseller models.Reseller
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&reseller)
	r:=models.CreateReseller(reseller)
	return r
}
func UpdateReseller(c *gin.Context)(models.Reseller){
	var reseller models.Reseller
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&reseller)
	r:=models.UpdateReseller(reseller)
	return r
}
func DeleteReseller(c *gin.Context)(models.Reseller){
	var reseller []models.Reseller
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&reseller)
	r:=models.DeleteReseller(reseller)
	return r
}
func GetReseller(c *gin.Context)([]models.Reseller){
	var reseller models.Reseller
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&reseller)
	r:=models.GetReseller(reseller)
	return r
}
func GetOwnerReseller(c *gin.Context)([] models.OwnerReseller){
	var owner models.OwnerReseller
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&owner)
	r:=models.GetOwnerReseller(owner)
	return r
}
func GetResellerForDatatable(c *gin.Context)(DatatableAjaxReseller){
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
        shortby="nama"
    case 2:
        shortby="idupline"
    }
    var price [] models.Reseller
    var total int64
    price,total=models.GetResellerForDatatable(search,start,length,shortby,typ,models.Reseller{Idupline:idcompany})

	ajax:=DatatableAjaxReseller{Draw:draw,TotalFiltered:len(price),Total:total,Data:price}
	return ajax
}