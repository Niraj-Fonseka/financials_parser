package controllers

import (
	"google_sheet_parser/App/models"

	"github.com/gin-gonic/gin"
)

func GETmonthlySummary(c *gin.Context) {
	month := c.Param("month")

	bymonth, err := models.MonthlySummary(month)

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
