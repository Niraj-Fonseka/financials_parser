package controllers

import (
	"github.com/gin-gonic/gin"
)

func GEThealth(c *gin.Context) {

	c.JSON(200, gin.H{
		"Health Status": "Up and Running",
	})
}
