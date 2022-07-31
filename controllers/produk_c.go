package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	_"log"
	"encoding/json"
	"errors"
	"strings"
)

type DatatableAjaxJenisProduk struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.JenisProduk `json:"data"`
}
type DatatableAjaxStatusProduk struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.StatusProduk `json:"data"`
}
type DatatableAjaxProduk struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Produk `json:"data"`
}
type GenProd struct {
	Idprefix int `json:"idprefix"`
	Idjenis int `json:"idjenis"`
	Nominal string `json:"nominal"`
	Nama string `json:"nama"`
	Alias string `json:"alias"`
}

func GenerateProduk(c *gin.Context)(models.Produk,error){
	var gp GenProd
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&gp)
	nominalArr:=strings.Split(gp.Nominal, ",")
	var p models.Produk
	if len(nominalArr)>0{
		var kodeArr[]string
		var namaArr[]string
		for i:=0; i<len(nominalArr); i++{
			kodeArr = append(kodeArr, strings.Replace(gp.Alias, "<DENOM>",nominalArr[i], 1))
			namaArr = append(namaArr, strings.Replace(gp.Nama, "<DENOM>",nominalArr[i], 1))
		}
		exist,err:=CheckAliasExist(gp.Idprefix,kodeArr)		
		exist2,err2:=CheckNamaExist(gp.Idprefix,namaArr)
		if exist{
			return models.Produk{},err
		}else if exist2{
			return models.Produk{},err2
		}else{
			for i:=0; i<len(nominalArr); i++{
				nom,_:=strconv.Atoi(nominalArr[i])
				p=models.Produk{Nama:namaArr[i],Alias:kodeArr[i],Idprefix:gp.Idprefix,Nominal:nom,Idjenis:gp.Idjenis,Idstatus:1}
				models.CreateProduk(models.Produk{Nama:namaArr[i],Alias:kodeArr[i],Idprefix:gp.Idprefix,Nominal:nom,Idjenis:gp.Idjenis,Idstatus:1})
			}
		}
	}
	return p,nil
}
func GetProdukForDatatable(c *gin.Context)(DatatableAjaxProduk){
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
        shortby="alias"
    case 4:
        shortby="nominal"
    }
    var pref [] models.Produk
    var total int64
    pref,total=models.GetProdukForDatatable(search,start,length,shortby,typ,models.Produk{})

	ajax:=DatatableAjaxProduk{Draw:draw,TotalFiltered:len(pref),Total:total,Data:pref}
	return ajax
}
func GetStatusProdukForDatatable(c *gin.Context)(DatatableAjaxStatusProduk){
	search:=c.Query("search[value]")
	idcompany,_:=strconv.Atoi(c.Query("idcompany"))
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
    }
    var pref [] models.StatusProduk
    var total int64
    pref,total=models.GetStatusProdukForDatatable(search,start,length,shortby,typ,models.StatusProduk{Idcompany:idcompany})

	ajax:=DatatableAjaxStatusProduk{Draw:draw,TotalFiltered:len(pref),Total:total,Data:pref}
	return ajax
}
func checkExistProduk(p models.Produk)(bool,error){
	exist:=false
	var err error
	if(len(models.GetProduk(models.Produk{Nama:p.Nama}))>0){
		exist=true
		err=errors.New("Name is already exist")
	}else if len(models.GetProduk(models.Produk{Alias:p.Alias}))>0 {
		exist=true
		err=errors.New("Alias is already exist")
	}
	return exist,err
}
func checkExistJenisProduk(p models.JenisProduk)(bool,error){
	exist:=false
	var err error
	if(len(models.GetJenisProduk(models.JenisProduk{Idcompany:p.Idcompany,Nama:p.Nama}))>0){
		exist=true
		err=errors.New("Name is already exist")
	}
	return exist,err
}
func checkExistStatusProduk(p models.StatusProduk)(bool,error){
	exist:=false
	var err error
	if(len(models.GetStatusProduk(models.StatusProduk{Idcompany:p.Idcompany,Nama:p.Nama}))>0){
		exist=true
		err=errors.New("Name is already exist")
	}
	return exist,err
}
func CreateProduk(c *gin.Context)(models.Produk,error){
	var prefix models.Produk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&prefix)
	exist,err:=checkExistProduk(prefix)
	if exist{
		return models.Produk{},err
	}else{
		return models.CreateProduk(prefix),nil
	}
}
func CreateJenisProduk(c *gin.Context)(models.JenisProduk,error){
	var jenis models.JenisProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&jenis)
	exist,err:=checkExistJenisProduk(jenis)
	if exist{
		return models.JenisProduk{},err
	}else{
		return models.CreateJenisProduk(jenis),nil
	}
}
func CreateStatusProduk(c *gin.Context)(models.StatusProduk,error){
	var status models.StatusProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&status)
	exist,err:=checkExistStatusProduk(status)
	if exist{
		return models.StatusProduk{},err
	}else{
		return models.CreateStatusProduk(status),nil
	}
}
func UpdateProduk(c *gin.Context)(models.Produk){
	var produk models.Produk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&produk)
	return models.UpdateProduk(produk)	
}
func UpdateJenisProduk(c *gin.Context)(models.JenisProduk){
	var jenis models.JenisProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&jenis)
	return models.UpdateJenisProduk(jenis)	
}
func UpdateStatusProduk(c *gin.Context)(models.StatusProduk){
	var status models.StatusProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&status)
	return models.UpdateStatusProduk(status)	
}
func GetJenisProduk(c *gin.Context)([]models.JenisProduk){
	var jenis models.JenisProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&jenis)
	return models.GetJenisProduk(jenis)
}
func GetStatusProduk(c *gin.Context)([]models.StatusProduk){
	var status models.StatusProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&status)
	return models.GetStatusProduk(status)
}
func GetProduk(c *gin.Context)([]models.Produk){
	var produk models.Produk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&produk)
	return models.GetProduk(produk)
}
func DeleteJenisProduk(c *gin.Context)(models.JenisProduk){
	var jenis [] models.JenisProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&jenis)
	return models.DeleteJenisProduk(jenis)	
}
func DeleteStatusProduk(c *gin.Context)(models.StatusProduk){
	var status [] models.StatusProduk
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&status)
	return models.DeleteStatusProduk(status)	
}
func GetJenisProdukForDatatable(c *gin.Context)(DatatableAjaxJenisProduk){
	search:=c.Query("search[value]")
	idcompany,_:=strconv.Atoi(c.Query("idcompany"))
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
    }
    var pref [] models.JenisProduk
    var total int64
    pref,total=models.GetJenisProdukForDatatable(search,start,length,shortby,typ,models.JenisProduk{Idcompany:idcompany})

	ajax:=DatatableAjaxJenisProduk{Draw:draw,TotalFiltered:len(pref),Total:total,Data:pref}
	return ajax
}