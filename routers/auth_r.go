package routers
import (
	"os"
	_"net/http"
	_"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"hris/controllers"
)
func AuthRouter(r *gin.Engine) {	
	r.LoadHTMLGlob("pages/*")
	r.Static("assets/", "./assets")

	r.GET("/", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "landing.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/login", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "login.html", gin.H{"appname":os.Getenv("APP_NAME"),"cid":os.Getenv("GoogleOuthCID")})
	})
	r.GET("/resetpassword", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "forgotpassword.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/register", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "register.html", gin.H{"appname":os.Getenv("APP_NAME"),"cid":os.Getenv("GoogleOuthCID"),"fid":os.Getenv("AppFbID")})
	})
	r.GET("/terms", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "terms.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.POST("/checkemail", func(c *gin.Context) { //Untuk Web
		c.JSON(200, controllers.CheckEmailExist(c))
	})
	r.POST("/checkusername", func(c *gin.Context) { //Untuk Web
		c.JSON(200, controllers.CheckUsernameExist(c))
	})
	r.POST("/trylogin", func(c *gin.Context) { //Untuk Web
		log.Println("trylogin")
		// if controllers.ClaimToken(c).Username ==""{
		// 	c.Redirect(http.StatusFound,"/login")
		// }else{
		// 	c.Redirect(http.StatusFound,"/")
		// }
	})
}