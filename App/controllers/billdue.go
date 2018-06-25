package controllers

import (
	"google_sheet_parser/App/models"

	"github.com/gin-gonic/gin"
)

func GETBillDueDates(c *gin.Context) {

	bymonth, err := models.BillDueDates()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"biil_due_dates": bymonth,
		})
	}
}
