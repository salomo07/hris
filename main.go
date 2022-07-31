package main
import (
	"os"
	"github.com/gin-gonic/gin"
	_"log"
  	"hris/routers"
)

func main() {
	r := gin.Default()
	port:=os.Getenv("PORT")
	// routers.ApiRouter(r)
	routers.AuthRouter(r)
	// routers.MasterRouter(r)
	// routers.AdministratorRouter(r)

	r.NoRoute(func(c *gin.Context) {
	    c.HTML(200, "404.html",gin.H{"appname":os.Getenv("APP_NAME")})
	})
	r.Run(":"+port)

	// port := os.Getenv("PORT")
	// r.Run(":"+port)
}
