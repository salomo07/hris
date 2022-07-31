package routers
import (
	"github.com/gin-gonic/gin"	
	"os"	
	"hris/controllers"
	"encoding/base64"
	_"log"
)
func MasterRouter(r *gin.Engine){
	r.Static("master/assets", "./assets")
	master := r.Group("/master")
	{
		master.GET("/user", func(c *gin.Context) {
			token,_ := c.Cookie("token")
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.HTML(200, "masteruser.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c),"token":token,"envidrole":os.Getenv("IDROLE_SU")})
			}else{
				c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
			}
		})
		master.GET("/application", func(c *gin.Context) {
			api:=controllers.CheckAPI(c)
			if api.Name != ""{
				c.HTML(200, "masterapplication.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c)})
			}
		})
		master.GET("/company", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != ""{
				c.HTML(200, "mastercompany.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c)})
			}
		})
		master.GET("/role", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != ""{
				c.HTML(200, "masterrole.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c),"envidrole":os.Getenv("IDROLE_SU")})
			}
		})
		master.GET("/menu", func(c *gin.Context) {
			dataSession:=controllers.CheckAPI(c)
			if dataSession.Name != ""{
				c.HTML(200, "mastermenu.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c)})
			}
		})
		master.GET("/access", func(c *gin.Context) {
			dataSession:=controllers.CheckAPI(c)
			if dataSession.Name != ""{
				c.HTML(200, "masteraccess.html", gin.H{"appname":os.Getenv("APP_NAME"),"datasession":controllers.ClaimToken(c),"envidrole":os.Getenv("IDROLE_SU")})
			}
		})


		master.GET("/getuser", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetUserForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		master.GET("/getapplication", func(c *gin.Context) {
			dataSession:=controllers.CheckAPI(c)
			if dataSession.Name != "" {
				c.JSON(200, controllers.GetApplicationForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		master.GET("/getcompanies", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetCompanyForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		master.GET("/getroles", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetRoleForDatatable(c))
			}
		})
		master.GET("/getmenus", func(c *gin.Context) {
			dataSession:=controllers.CheckAPI(c)
			if dataSession.Name != "" {
				c.JSON(200, controllers.GetMenuForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		master.GET("/getsubmenus", func(c *gin.Context) {
			dataSession:=controllers.CheckAPI(c)
			if dataSession.Name != "" {
				c.JSON(200, controllers.GetSubmenuForDatatable(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
	}
	masterApi := r.Group("/api/master")
	{	
		masterApi.POST("/getcompany", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetCompany(c))
			}
		})
		masterApi.POST("/createcompany", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.CreateCompany(c)
				if res.Name!=""{
					c.JSON(200,res)
				}				
			}
		})
		masterApi.POST("/updatecompany", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.UpdateCompany(c)
				if res.Name!=""{
					c.JSON(200, res)
				}
			}
		})
		masterApi.POST("/deletecompany", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.DeleteCompany(c)
				if res.Name!=""{
					c.JSON(200, res)
				}
			}
		})

		masterApi.POST("/getrole", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetRole(c))
			}
		})
		masterApi.POST("/createrole", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.CreateRole(c)
				if res.Role!=""{
					c.JSON(200,res)
				}
			}
		})
		masterApi.POST("/updaterole", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.UpdateRole(c)
				if res.Role!=""{
					c.JSON(200, res)
				}
			}
		})
		masterApi.POST("/deleterole", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				res:=controllers.DeleteRole(c)
				if res.Role!=""{
					c.JSON(200, res)
				}
			}
		})

		masterApi.POST("/getuser", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetUser(c))
			}
		})
		masterApi.POST("/createuser", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				data:=controllers.CreateUser(c)
				if data.Username!=""{
					c.JSON(200, data)
				}
			}
		})
		masterApi.POST("/updateuser", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				data:=controllers.UpdateUser(c)
				if data.Username!=""{
					c.JSON(200, data)
				}
			}
		})
		masterApi.POST("/deleteuser", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.DeleteUser(c))
			}
		})





		masterApi.GET("/getallsubMenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetSubmenu())
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.GET("/getallmenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetMenu())
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/createaccess1", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200,controllers.CreateAccess1(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/createaccess2", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200,controllers.CreateAccess2(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/updateaccess1", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.UpdateAccess1(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/updateaccess2", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.UpdateAccess2(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/createsubmenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.CreateSubmenu(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/createmenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.CreateMenu(c))
			}
		})
		masterApi.POST("/updatemenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.UpdateMenu(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/updatesubmenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.UpdateSubmenu(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/deletemenu", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.DeleteMenu(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		
		masterApi.POST("/createapplication", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.CreateApplication(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/updateapplication", func(c *gin.Context){
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.UpdateApplication(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})

		masterApi.POST("/generateapi", func(c *gin.Context){
			api:=controllers.CheckAPI(c)
			if api.Name != "" {
				c.JSON(200, controllers.GenerateAPIKey(c))
			}
		})
		
		masterApi.POST("/getaccess1", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetAccess1(c))
			}
		})
		masterApi.POST("/getaccess2", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.GetAccess2(c))
			}
		})
		masterApi.POST("/givefullaccess", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				controllers.GiveFullAccess(c)
				c.JSON(200,"{result:'All access given'}" )
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/generateuserownerexist", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				
				c.JSON(200,controllers.GenerateUserOwnerExist(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		
		masterApi.POST("/addAllAccess1", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.AddAllAccess1(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
		masterApi.POST("/addAllAccess2", func(c *gin.Context) {
			app:=controllers.CheckAPI(c)
			if app.Name != "" {
				c.JSON(200, controllers.AddAllAccess2(c))
			}else{
				c.JSON(401, gin.H{"result":"Unauthorized"})
			}
		})
	}
}