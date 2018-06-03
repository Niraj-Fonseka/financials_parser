package controllers

import (
	"google_sheet_parser/App/models"

	"github.com/gin-gonic/gin"
)

func GETbillsList(c *gin.Context) {

	billsList := models.DashboardBillsList()
	c.JSON(200, gin.H{
		"bills": billsList,
	})

}

func GETbillsListMonthly(c *gin.Context) {
	month := c.Param("month")

	bymonth, err := models.DashboardBillsListMonthly(month)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"bymonth": bymonth,
		})
	}

}
