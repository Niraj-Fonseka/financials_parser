package routers

import (
	"google_sheet_parser/App/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	//HealthCheck
	router.GET("/health", controllers.GEThealth)

	//Data in the dashboard page
	dash := router.Group("/dashboard")
	{
		dash.GET("/bills", controllers.GETbillsList)
	}

	return router
}
