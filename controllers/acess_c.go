package controllers
import (
	"github.com/gin-gonic/gin"	
	"hris/models"
	"strconv"
	"log"
	"encoding/json"
)
type DatatableAjaxAccess1 struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Access1 `json:"data"`
}
type DatatableAjaxAccess2 struct {
	Draw int `json:"draw"`
	Total int64 `json:"recordsTotal"`
	TotalFiltered int `json:"recordsFiltered"`
	Data []models.Access2 `json:"data"`
}

func AddAllAccess1(c *gin.Context)(models.Access1){
	var acc1 models.Access1
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc1)
	rA:=models.GetAccess1(acc1)
	rM:=models.GetMenu(models.Menu{})

	for i := 0; i < len(rM); i++ {
	    log.Println(rM[i],rA)
	    if len(rA)==0{
	    	models.CreateAccess1(models.Access1{Idrole:acc1.Idrole,Idmenu1:rM[i].Id,Create:true,Read:true,Update:true,Delete:true})
	    }else{
	    	exist:=false
	    	for j := 0; j < len(rA); j++ {
	    		if rA[j].Idmenu1==rM[i].Id{
	    			exist=true
	    		}
	    	}
	    	if exist==false {
	    		models.CreateAccess1(models.Access1{Idrole:acc1.Idrole,Idmenu1:rM[i].Id,Create:true,Read:true,Update:true,Delete:true})
	    	}
	    }
	}
	return acc1
}
func AddAllAccess2(c *gin.Context)(models.Access2){
	var acc2 models.Access2
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc2)
	rA:=models.GetAccess2(acc2)
	rM:=models.GetSubmenu(models.Submenu{})


	for i := 0; i < len(rM); i++ {
	    if len(rA)==0{
	    	models.CreateAccess2(models.Access2{Idrole:acc2.Idrole,Idmenu2:rM[i].Id,Create:true,Read:true,Update:true,Delete:true})
	    }else{
	    	exist:=false
	    	for j := 0; j < len(rA); j++ {
	    		if rA[j].Idmenu2==rM[i].Id{
	    			exist=true
	    		}
	    	}
	    	if exist==false {
	    		models.CreateAccess2(models.Access2{Idrole:acc2.Idrole,Idmenu2:rM[i].Id,Create:true,Read:true,Update:true,Delete:true})
	    	}
	    }
	}
	return acc2
}
func GetAccess2(c *gin.Context)([]models.Access2){
	var acc2 models.Access2
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc2)
	user:=ClaimToken(c)
	acc2.Idapplication=user.Idapplication
	r:=models.GetAccess2(acc2)
	return r
}
func GetAccess1(c *gin.Context)([]models.Access1){
	var acc1 models.Access1
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc1)
	user:=ClaimToken(c)
	acc1.Idapplication=user.Idapplication
	r:=models.GetAccess1(acc1)
	return r
}
func CreateAccess1(c *gin.Context)(models.Access1){
	var acc models.Access1
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc)
	r:=models.CreateAccess1(acc)
	return r
}
func CreateAccess2(c *gin.Context)(models.Access2){
	var acc2 models.Access2
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc2)
	r:=models.CreateAccess2(acc2)
	return r
}
func UpdateAccess1(c *gin.Context)(models.Access1){
	var acc models.Access1
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc)
	r:=models.UpdateAccess1(acc)
	return r
}
func UpdateAccess2(c *gin.Context)(models.Access2){
	var acc2 models.Access2
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&acc2)
	r:=models.UpdateAccess2(acc2)
	return r
}
// func DeleteAccess1(c *gin.Context)(models.Access1){
// 	var acc models.Access1
// 	jsonData, _ := c.GetRawData()
// 	json.Unmarshal(jsonData,&acc)
// 	acc.Active=false
// 	r:=models.DeleteAccess1(acc)
// 	return r
// }
// func DeleteAccess2(c *gin.Context)(models.Access2){
// 	var acc2 models.Access2
// 	jsonData, _ := c.GetRawData()
// 	json.Unmarshal(jsonData,&acc2)
// 	acc2.Active=false
// 	r:=models.DeleteAccess2(acc2)
// 	return r
// }
func GetAccess1ForDatatable(c *gin.Context)(DatatableAjaxAccess1){
	search:=c.Query("search[value]")
	start,_:=strconv.Atoi(c.Query("start"))
	length,_:=strconv.Atoi(c.Query("length"))
	draw,_:=strconv.Atoi(c.Query("draw"))
	app:=CheckAPI(c)
    var acc [] models.Access1
    var total int64
    acc,total=models.GetAccess1ForDatatable(search,start,length,models.Access1{Idapplication:app.Id})

	ajax:=DatatableAjaxAccess1{Draw:draw,TotalFiltered:len(acc),Total:total,Data:acc}

	log.Println(ajax,"dfghz")
	return ajax
}
func GetAccess2ForDatatable(c *gin.Context)(DatatableAjaxAccess2){
	search:=c.Query("search[value]")
	start,_:=strconv.Atoi(c.Query("start"))
	length,_:=strconv.Atoi(c.Query("length"))
	draw,_:=strconv.Atoi(c.Query("draw"))
	
	app:=CheckAPI(c)
    var acc [] models.Access2
    var total int64
    acc,total=models.GetAccess2ForDatatable(search,start,length,models.Access2{Idapplication:app.Id})
	ajax:=DatatableAjaxAccess2{Draw:draw,TotalFiltered:len(acc),Total:total,Data:acc}
	return ajax
}
func GiveFullAccess(c *gin.Context){
	var role models.Role	
	jsonData, _ := c.GetRawData()
	json.Unmarshal([]byte(jsonData),&role)
	menu1:=models.GetMenu1CantAccess(role)
	menu2:=models.GetMenu2CantAccess(role)
	for i := 0; i < len(menu1); i++ {
	    models.CreateAccess1(models.Access1{Create:true,Read:true,Update:true,Delete:true,Idmenu1:menu1[i].Id,Idrole:role.Id})
	}
	for i := 0; i < len(menu2); i++ {
	    models.CreateAccess2(models.Access2{Create:true,Read:true,Update:true,Delete:true,Idmenu2:menu2[i].Id,Idrole:role.Id})
	}
}