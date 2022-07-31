 package routers
import (
	"github.com/gin-gonic/gin"	
	_"os"	
	"hris/controllers"
)
func RoleRouter(r *gin.Engine){
	r.Static("role/assets", "./assets")
	role := r.Group("/role")
	{
		role.POST("/getallrole", func(c *gin.Context) {
			if controllers.CheckAPI(c).Name != ""{
				c.JSON(200, controllers.GetUserForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
	}
}