package main

import (
	"google_sheet_parser/App/models"
	"google_sheet_parser/App/routers"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true

	// Register database
	// log.Println("Registering database..")
	// models.OpenDB(os.Getenv("DBCONN"), false) // Pass false for not logging dababase queries
	// defer models.CloseDB(models.DB)

	log.SetOutput(gin.DefaultWriter)             // logs gin
	log.SetFlags(log.LstdFlags | log.Lshortfile) // logs with timestamp and file name and line number

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	models.InitGoogleSheetClient()
	router = routers.InitRouters(router)

	//log.Println("AppName is: ", config.Get("server.appName").(string))

	router.Run(":7070")
}
