package routers
import (
	"os"
	"net/http"
	_"encoding/base64"
	"github.com/gin-gonic/gin"
	_"log"
	"hris/controllers"
)
func AuthRouter(r *gin.Engine) {	
	r.LoadHTMLGlob("pages/*")
	r.Static("assets/", "./assets")

	r.GET("/", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "index.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/login", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "login.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.POST("/trysignin", func(c *gin.Context) { //Untuk Web
		if controllers.ClaimToken(c).Username ==""{
			c.Redirect(http.StatusFound,"/login")
		}else{
			c.Redirect(http.StatusFound,"/")
		}
	})
	// r.GET("/login", func(c *gin.Context) { //Untuk Web
	// 	c.HTML(200, "login.html", gin.H{"appname":os.Getenv("APP_NAME")})
		
	// })
}