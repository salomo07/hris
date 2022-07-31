package routers
import (
	"github.com/gin-gonic/gin"
	_"log"
)
func AdministratorRouter(r *gin.Engine) {
	r.Static("administrator/assets", "./assets")	
	administrator := r.Group("/administrator")
	{
		ProdukWebRouter(administrator)
		PriceWebRouter(administrator)
		PrefixWebRouter(administrator)
		ResellerWebRouter(administrator)
	}
	// API
	administratorApi := r.Group("/api/administrator")
	{	
		ProdukAPIRouter(administratorApi)
		PriceAPIRouter(administratorApi)
		PrefixAPIRouter(administratorApi)
		ResellerAPIRouter(administratorApi)
	}
}
