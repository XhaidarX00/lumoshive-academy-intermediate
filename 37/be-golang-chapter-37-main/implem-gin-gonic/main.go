package main

import (
	"golang-chapter-37/implem-gin-gonic/config"
	// "golang-chapter-37/implem-gin-gonic/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LoadConfig()
}

func main() {
	r := gin.Default()

	// Middleware
	// r.Use(middleware.Logger())
	// r.Use(middleware.BasicAuth())

	// db := database.GetDB()
	// sc := repository.NewShippingRepository()
	// controller := controller.NewController()
	// router.APIRouter(r, controller)

	r.Run(viper.GetString("SERVER_PORT"))
}
