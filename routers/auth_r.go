package routers
import (
	"os"
	_"net/http"
	_"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	_"hris/controllers"
)
func AuthRouter(r *gin.Engine) {	
	r.LoadHTMLGlob("pages/*")
	r.Static("assets/", "./assets")

	r.GET("/", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "landing.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/login", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "login.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/resetpassword", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "forgotpassword.html", gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.GET("/register", func(c *gin.Context) { //Untuk Web
		c.HTML(200, "register.html", gin.H{"appname":os.Getenv("APP_NAME")})
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