package main

import (
	"log"
	_ "project-voucher-team3/docs"
	"project-voucher-team3/infra"
	"project-voucher-team3/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server for a Swagger API.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url https://e-commers-darmi/
// @contact.email darmi.ecommers@gmail.com
// @license.name E-Commers Darmi
// @license.url https://darmi.ecommers.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey IDKeyAuth
// @in header
// @name IDKey
func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
