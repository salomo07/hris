 package routers
import (
	"github.com/gin-gonic/gin"
	"os"
	"encoding/base64"
	"hris/controllers"
)
func ProdukWebRouter(administrator *gin.RouterGroup){
	administrator.GET("/produk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != ""{
			
			c.HTML(200, "adminproduk.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":dataSession,"operator":controllers.GetPrefixForDatatable(c)})
		}else{
			c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
		}
	})
	administrator.GET("/getjenisproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetJenisProdukForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administrator.GET("/getstatusproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetStatusProdukForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administrator.GET("/getproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetProdukForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}
func ProdukAPIRouter(administratorApi *gin.RouterGroup){
	administratorApi.POST("/generateproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.GenerateProduk(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}					
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/getjenisproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetJenisProduk(c))					
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/createstatusproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.CreateStatusProduk(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})	
	administratorApi.POST("/createjenisproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.CreateJenisProduk(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/createproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.CreateProduk(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}
			
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/updateproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.UpdateProduk(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/updatejenisproduk", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.UpdateJenisProduk(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}