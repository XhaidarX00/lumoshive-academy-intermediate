// routes/routes.go
package routes

import (
	"project_auth_jwt/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/report/:minute", ctx.Ctl.Report.GenerateReport)
	authRoutes(r, ctx)
	userRoutes(r, ctx)
	return r
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	publicRoutes := r.Group("/")
	// publicRoutes.Use(ctx.Middleware.Limit.IPWhitelistMiddleware())
	{
		publicRoutes.POST("/register", ctx.Ctl.Auth.Register)
		publicRoutes.POST("/login", ctx.Middleware.Limit.Limit(), ctx.Ctl.Auth.Login)
		// publicRoutes.POST("/login", ctx.Ctl.Auth.Login)
	}
}

func userRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(ctx.Middleware.Limit.IPWhitelistMiddleware(), ctx.Middleware.Auth.Authenticate())
	{
		protectedRoutes.GET("/users", ctx.Ctl.User.GetAllUsers)
		protectedRoutes.GET("/users/:id", ctx.Ctl.User.GetUserByID)
	}
}
