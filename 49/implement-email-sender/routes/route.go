package routes

import (
	"dashboard-ecommerce-team2/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := ctx.Middleware.Authentication()
	adminMiddleware := ctx.Middleware.RoleAuthorization("admin")

	productRoutes := r.Group("/products", authMiddleware)
	{
		productRoutes.POST("/", ctx.Ctl.Product.CreateProductController)
		productRoutes.GET("/", ctx.Ctl.Product.GetAllProductsController)
		productRoutes.GET("/:id", ctx.Ctl.Product.GetProductByIDController)
		productRoutes.DELETE("/:id", adminMiddleware, ctx.Ctl.Product.DeleteProductController)
		productRoutes.PUT("/:id", ctx.Ctl.Product.UpdateProductController)
	}
	stockRoutes := r.Group("/stock", authMiddleware)
	{
		stockRoutes.GET("/:id", ctx.Ctl.Stock.GetProductStockDetailController)
		stockRoutes.DELETE("/:id", adminMiddleware, ctx.Ctl.Stock.DeleteProductStockController)
		stockRoutes.PUT("/", ctx.Ctl.Stock.UpdateProductStockController)
	}

	orderRoutes := r.Group("/orders", authMiddleware)
	{
		orderRoutes.GET("/", ctx.Ctl.Order.GetAllOrdersController)
		orderRoutes.GET("/:id", ctx.Ctl.Order.GetOrderByIDController)
		orderRoutes.PUT("/update/:id", ctx.Ctl.Order.UpdateOrderStatusController)
		orderRoutes.DELETE("/:id", adminMiddleware, ctx.Ctl.Order.DeleteOrderController)
		orderRoutes.GET("/detail/:id", ctx.Ctl.Order.GetOrderDetailController)
	}

	categoryRoutes := r.Group("/category", adminMiddleware)
	{
		categoryRoutes.POST("/create", ctx.Ctl.Category.CreateCatergoryController)
		categoryRoutes.GET("/list", ctx.Ctl.Category.GetAllCategoriesController)
		categoryRoutes.GET("/:id", ctx.Ctl.Category.GetCategoryByIDController)
		categoryRoutes.PUT("/update/:id", ctx.Ctl.Category.UpdateCategoryController)
		categoryRoutes.DELETE("/:id", adminMiddleware, ctx.Ctl.Category.DeleteCategoryController)
	}

	authRoutes(r, ctx)
	dashboardRoutes(r, ctx)
	bannerRoutes(r, ctx)
	promotionRoutes(r, ctx)
	return r
}

func bannerRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	banner := r.Group("/api", ctx.Middleware.Authentication())
	banner.GET("/banner", ctx.Ctl.Banner.GetBannerByIDController)
	banner.PUT("/banner", ctx.Ctl.Banner.UpdateBannerController)
	banner.POST("/create-banner", ctx.Ctl.Banner.CreateBannerController)
	banner.DELETE("/banner", ctx.Middleware.RoleAuthorization("admin"), ctx.Ctl.Banner.DeleteBannerController)
}

func promotionRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	promotion := r.Group("/api", ctx.Middleware.Authentication())
	promotion.GET("/list-promotion", ctx.Ctl.Promotion.GetAllPromotionsController)
	promotion.GET("/promotion", ctx.Ctl.Promotion.GetByIdPromotionsController)
	promotion.PUT("/promotion", ctx.Ctl.Promotion.UpdatePromotionController)
	promotion.POST("/create-promotion", ctx.Ctl.Promotion.CreatePromotionController)
	promotion.DELETE("/promotion", ctx.Middleware.RoleAuthorization("admin"), ctx.Ctl.Promotion.DeletePromotionController)
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", ctx.Ctl.User.LoginController)
	authGroup.POST("/check-email", ctx.Ctl.User.CheckEmailUserController)
	authGroup.POST("/register", ctx.Ctl.User.CreateUserController)
	authGroup.GET("/validate-otp", ctx.Ctl.User.ValidateOTPHandler)
	authGroup.GET("/send-otp", ctx.Ctl.User.ResetPassOtpSenderController)
	authGroup.PATCH("/reset-password", ctx.Ctl.User.ResetUserPasswordController)
}

func dashboardRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(ctx.Middleware.Authentication())
	{
		dashboardGroup.GET("/summary", ctx.Ctl.Dashboard.GetSummaryController)
		dashboardGroup.GET("/current-month-earning", ctx.Ctl.Dashboard.CurrentMonthEarningController)
		dashboardGroup.GET("/revenue-chart", ctx.Ctl.Dashboard.RenevueChartController)
		dashboardGroup.GET("/best-item-list", ctx.Ctl.Dashboard.GetBestProductListController)
	}
}
