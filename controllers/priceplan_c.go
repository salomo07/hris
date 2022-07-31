package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	_"log"
	"encoding/json"
	"errors"
)
func GetPricePlanForDatatable(c *gin.Context)(DatatableAjaxPricePlan){
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
    }
    user:=ClaimToken(c)
    var price [] models.PricePlan
    var total int64
    price,total=models.GetPricePlanForDatatable(search,start,length,shortby,typ,models.PricePlan{Idcompany:user.Idcompany})
	ajax:=DatatableAjaxPricePlan{Draw:draw,TotalFiltered:len(price),Total:total,Data:price}
	return ajax
}
func GetPricePlan(price models.PricePlan)([] models.PricePlan){
	return models.GetPricePlan(price)
}
func GetPricePlanDetailForDatatable(c *gin.Context)(DatatableAjaxPricePlanDetail){
	idpriceplan,_:=strconv.Atoi(c.Query("idpriceplan"))
	draw,_:=strconv.Atoi(c.Query("draw"))
    var price [] models.PricePlanDetail
    var total int64
    price,total=models.GetPricePlanDetailForDatatable(idpriceplan)

	ajax:=DatatableAjaxPricePlanDetail{Draw:draw,TotalFiltered:len(price),Total:total,Data:price}
	return ajax
}
func checkExistPricePlan(p models.PricePlan)(bool,error){
	exist:=false
	var err error
	if(len(models.GetPricePlan(models.PricePlan{Nama:p.Nama}))>0){
		exist=true
		err=errors.New("Name is already exist")
	}
	return exist,err
}
func CreatePricePlan(c *gin.Context)(models.PricePlan,error){
	var price models.PricePlan
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&price)
	exist,err:=checkExistPricePlan(price)
	if exist{
		return models.PricePlan{},err
	}else{
		return models.CreatePricePlan(price),nil
	}
}
func UpsertPricePlanDetail(c *gin.Context)([]models.PricePlanDetail,error){
	var prices []models.PricePlanDetail
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prices)
	for i:=0;i<len(prices);i++{
		models.UpsertPricePlanDetail(prices[i])
	}
	return prices,nil
}
func UpdatePricePlan(c *gin.Context)(models.PricePlan){
	var price models.PricePlan
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&price)
	return models.UpdatePricePlan(price)	
}
func DeletePricePlan(c *gin.Context)(models.PricePlan){
	var price [] models.PricePlan
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&price)
	return models.DeletePricePlan(price)	
}