 package routers
import (
	"github.com/gin-gonic/gin"
	"os"
	"encoding/base64"
	"hris/controllers"
)
func PriceWebRouter(administrator *gin.RouterGroup){
	administrator.GET("/price", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != ""{			
			c.HTML(200, "adminprice.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":dataSession,"operator":controllers.GetPrefixForDatatable(c),"dataproduk":controllers.GetProdukForDatatable(c)})
		}else{
			c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
		}
	})
	administrator.GET("/getpriceplan", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetPricePlanForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}
func PriceAPIRouter(administratorApi *gin.RouterGroup){
	administratorApi.GET("/getdetailpriceplan", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.GetPricePlanDetailForDatatable(c))				
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/createpriceplan", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.CreatePricePlan(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/updatepriceplan", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.UpdatePricePlan(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/deletepriceplan", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.DeletePricePlan(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/price/upsertPricePlanDetail", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.UpsertPricePlanDetail(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}					
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}