package controllers

import (
	"google_sheet_parser/App/models"

	"github.com/gin-gonic/gin"
)

func GETmonthlySpendingSummary(c *gin.Context) {
	month := c.Param("month")

	bymonth, err := models.MonthlySpeditures(month)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"spending": bymonth,
		})
	}

}

func GETmonthlySummary(c *gin.Context) {
	month := c.Param("month")

	summary, err := models.DashboardMonthlySummary(month)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"summary": summary,
		})
	}

}
