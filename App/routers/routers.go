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
		dash.GET("/bills/:month", controllers.GETbillsListMonthly)
		dash.GET("/summary/:month", controllers.GETmonthlySummary)

	}

	//Data in month pages
	monthly := router.Group("/monthly")
	{
		monthly.GET("/spending/:month", controllers.GETmonthlySpendingSummary)
	}

	//Data in bill due date page
	duedates := router.Group("/bills")
	{
		duedates.GET("/duedates", controllers.GETBillDueDates)
	}
	return router
}
