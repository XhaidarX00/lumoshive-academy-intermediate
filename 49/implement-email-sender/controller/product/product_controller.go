package productcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewProductController(service service.Service, log *zap.Logger) *ProductController {
	return &ProductController{
		Service: service,
		Log:     log,
	}
}

// CreateProductController godoc
// @Summary Creates a new product
// @Description Create a new product with an image
// @Tags Product
// @Accept json
// @Produce json
// @Param category_id formData string true "Category ID"
// @Param name formData string true "Product Name"
// @Param code_product formData string true "Code Product"
// @Param description formData string true "Description Product"
// @Param price formData number true "Product Price"
// @Param stock formData int true "Product Stock"
// @Param image formData file true "Product Image"
// @Success 201 {object} utils.ResponseOK{data=models.Product} "Product created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid input"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Security Authentication
// @Security UserID
// @Router /products [post]
func (ctrl *ProductController) CreateProductController(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBind(&input); err != nil {
		ctrl.Log.Error("handler: Failed to bind input", zap.Error(err))
		helper.ResponseError(c, "invalid payload", helper.FormatValidationError(err), http.StatusBadRequest)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		ctrl.Log.Debug("handler: Failed to get uploaded file", zap.Error(err))
		ctrl.Log.Error("handler: Failed to get uploaded file", zap.Error(err))
		helper.ResponseError(c, "image is required", err.Error(), http.StatusBadRequest)
		return
	}

	filePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		ctrl.Log.Debug("handler: Failed to save uploaded file", zap.Error(err))
		ctrl.Log.Error("handler: Failed to save uploaded file", zap.Error(err))
		helper.ResponseError(c, "Failed to save file", err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.Log.Info("File saved successfully", zap.String("file_path", filePath))

	product, err := ctrl.Service.Product.CreateProduct(&input, filePath)
	if err != nil {
		ctrl.Log.Debug("handler: Failed to create product", zap.Error(err))
		ctrl.Log.Error("handler: Failed to create product", zap.Error(err))
		helper.ResponseError(c, "Failed to create product", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.Log.Info("handler: Product created successfully", zap.Int("product_id", product.ID))
	helper.ResponseOK(c, product, "create successfully", http.StatusCreated)

}

// GetAllProductsController godoc
// @Summary Get all products with pagination
// @Description Get a paginated list of all products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginationResponse{Data=[]models.ProductWithCategory} "List of products"
// @Failure 500 {object} utils.ErrorResponse "Failed to fetch products"
// @Security Authentication
// @Security UserID
// @Router /products [get]
func (ctrl *ProductController) GetAllProductsController(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt := helper.StringToInt(page)
	pageSizeInt := helper.StringToInt(pageSize)

	products, totalItems, err := ctrl.Service.Product.GetAllProducts(pageInt, pageSizeInt)
	if err != nil {
		if err.Error() == "Product not found" {
			ctrl.Log.Warn("handler: No products found", zap.Int("page", pageInt), zap.Int("pageSize", pageSizeInt))
			helper.ResponseError(c, "No products found", err.Error(), http.StatusNotFound)
			return
		}

		ctrl.Log.Debug("handler: Failed to fetch products", zap.Error(err))
		ctrl.Log.Error("handler: Failed to fetch products", zap.Error(err))
		helper.ResponseError(c, "Failed to fetch products", err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := (totalItems + pageSizeInt - 1) / pageSizeInt

	helper.ResponseOKPagination(c, products, "", pageInt, pageSizeInt, totalItems, totalPages, http.StatusOK)
}

// GetProductByIDController godoc
// @Summary Get product by ID
// @Description Get a specific product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.ResponseOK{data=models.Product} "Product details"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Security Authentication
// @Security UserID
// @Router /products/{id} [get]
func (ctrl *ProductController) GetProductByIDController(c *gin.Context) {
	id := c.Param("id")

	productID := helper.StringToInt(id)

	product, err := ctrl.Service.Product.GetProductByID(productID)
	if err != nil {
		ctrl.Log.Warn("handler: No products found", zap.Int("id", productID))
		helper.ResponseError(c, "No products found", err.Error(), http.StatusNotFound)
		return
	}

	helper.ResponseOK(c, product, "", http.StatusOK)

}

// UpdateProductController godoc
// @Summary Update a product by ID
// @Description Update an existing product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param category_id formData string false "Category ID"
// @Param name formData string false "Product Name"
// @Param code_product formData string false "Code Product"
// @Param description formData string false "Description Product"
// @Param price formData number false "Product Price"
// @Param stock formData int false "Product Stock"
// @Param image formData file false "Product Image"
// @Param published formData boolean false "Is Published"
// @Success 200 {object} utils.ResponseOK{data=models.Product} "Product updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid input"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Security Authentication
// @Security UserID
// @Router /products/{id} [put]
func (ctrl *ProductController) UpdateProductController(c *gin.Context) {
	id := c.Param("id")

	var input models.Product
	if err := c.ShouldBind(&input); err != nil {
		ctrl.Log.Error("handler: Failed to bind input", zap.Error(err))
		helper.ResponseError(c, "Invalid payload", helper.FormatValidationError(err), http.StatusBadRequest)
		return
	}

	file, err := c.FormFile("image")
	filePath := ""
	if err == nil {
		filePath = "/tmp/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			ctrl.Log.Debug("handler: Failed to save uploaded file", zap.Error(err))
			ctrl.Log.Error("handler: Failed to save uploaded file", zap.Error(err))
			helper.ResponseError(c, "Failed to save file", err.Error(), http.StatusInternalServerError)
			return
		}
		ctrl.Log.Info("File saved successfully", zap.String("file_path", filePath))
	}

	product, err := ctrl.Service.Product.UpdateProduct(helper.StringToInt(id), input, filePath)
	if err != nil {
		ctrl.Log.Debug("handler: Failed to update product", zap.Error(err))
		ctrl.Log.Error("handler: Failed to update product", zap.Error(err))
		helper.ResponseError(c, "Failed to update product", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.Log.Info("handler: Product updated successfully", zap.Int("product_id", product.ID))
	helper.ResponseOK(c, product, "Update successfully", http.StatusOK)
}

// DeleteProductController godoc
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.ResponseOK "Product deleted successfully"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Security Authentication
// @Security UserID
// @Security UserRole
// @Router /products/{id} [delete]
func (ctrl *ProductController) DeleteProductController(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))

	err := ctrl.Service.Product.DeleteProduct(id)
	if err != nil {
		if err.Error() == "product not found" {
			ctrl.Log.Warn("handler: No products found", zap.Int("id", id))
			helper.ResponseError(c, "No products found", err.Error(), http.StatusNotFound)
		} else {
			ctrl.Log.Debug("handler: Failed to delete product", zap.Int("id", id), zap.Error(err))
			ctrl.Log.Error("handler: Failed to delete product", zap.Int("id", id), zap.Error(err))
			helper.ResponseError(c, "Failed", err.Error(), http.StatusInternalServerError)
		}
		return
	}
	helper.ResponseOK(c, nil, "Product deleted successfully", http.StatusOK)
}
