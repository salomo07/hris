 package routers
import (
	"github.com/gin-gonic/gin"	
	"os"
	_"log"
	_"strconv"
	_"encoding/base64"
	"hris/models"	
	"hris/controllers"
)
func ResellerWebRouter(administrator *gin.RouterGroup){
	administrator.GET("/reseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != ""{
			c.HTML(200, "adminreseller.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c),
			"datapriceplan":controllers.GetPricePlan(models.PricePlan{Idcompany:1})})
		}
	})
	administrator.GET("/getreseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetResellerForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}
func ResellerAPIRouter(administratorAPI *gin.RouterGroup){
	administratorAPI.POST("/createreseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p:=controllers.CreateReseller(c)
			c.JSON(200,p)
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorAPI.POST("/updatereseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.UpdateReseller(c))
		}
	})
	administratorAPI.POST("/deletereseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.DeleteReseller(c))
		}
	})
	administratorAPI.POST("/getresellerbyid", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.GetReseller(c))
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorAPI.POST("/checkuserreseller", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.GetOwnerReseller(c))
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}