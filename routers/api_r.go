package routers
import (
	"github.com/gin-gonic/gin"
	_"log"
	"hris/controllers"
)
func ApiRouter(r *gin.Engine) {	
	api := r.Group("/api/auth")
	{
		api.POST("/login", func(c *gin.Context) {
			data:=controllers.CheckAPI(c)
			if data.Name!=""{
				c.JSON(200, gin.H{"data":data})
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		api.POST("/resetpassword", func(c *gin.Context) {
			data:=controllers.CheckAPI(c)
			if data.Name!=""{
				c.JSON(200, gin.H{"data":data})
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		
		// api.POST("/checkuserexist", func(c *gin.Context) {
		// 	data:=controllers.CheckAPI(c)
		// 	if data.Name!=""{			
		// 		c.JSON(200, gin.H{"data":controllers.GetUserByJson(c)})			
		// 	}else{
		// 		c.JSON(401, gin.H{"result":"Unauthorized"})
		// 	}
		// })
		// api.POST("/register", func(c *gin.Context) {
		// 	data:=controllers.CheckAPIAuthentication(c)
		// 	if data.Username!=""{			
		// 		c.JSON(200, gin.H{"data":controllers.CreateUser(c)})			
		// 	}else{
		// 		c.JSON(401, gin.H{"result":"Unauthorized"})
		// 	}
		// })
		// api.POST("/getRole", func(c *gin.Context) {
		// 	data:=controllers.CheckAPIAuthentication(c)
		// 	if data.Username!=""{			
		// 		c.JSON(200, gin.H{"data":controllers.CreateUser(c)})			
		// 	}else{
		// 		c.JSON(401, gin.H{"result":"Unauthorized"})
		// 	}
		// })
	}
	
}