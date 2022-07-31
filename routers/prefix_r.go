 package routers
import (
	"github.com/gin-gonic/gin"
	"os"
	"encoding/base64"
	"hris/controllers"
)
func PrefixWebRouter(administrator *gin.RouterGroup){
	administrator.GET("/prefix", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != ""{
			c.HTML(200, "adminprefix.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":dataSession,"envidrole":os.Getenv("IDROLE_SU")})
		}else{
			c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
		}
	})
	administrator.GET("/getprefix", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200, controllers.GetPrefixForDatatable(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}
func PrefixAPIRouter(administratorApi *gin.RouterGroup){
	administratorApi.POST("/createprefix", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			p,err:=controllers.CreatePrefix(c)
			if err !=nil {
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,p)
			}
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	administratorApi.POST("/updateprefix", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.UpdatePrefix(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
	
	administratorApi.POST("/deleteprefix", func(c *gin.Context) {
		dataSession:=controllers.CheckAPI(c)
		if dataSession.Name != "" {
			c.JSON(200,controllers.DeleteJenisProduk(c))
		}else{
			c.JSON(401, gin.H{"result":"Unauthorized"})
		}
	})
}