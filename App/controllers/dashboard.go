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
