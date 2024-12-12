package categorycontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	utils "dashboard-ecommerce-team2/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewCategoryController(service service.Service, log *zap.Logger) *CategoryController {
	return &CategoryController{
		Service: service,
		Log:     log,
	}
}

// CreateCategoryController godoc
// @Summary Create a new category
// @Description Create a new category with a name
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category name"
// @Success 201 {object} helper.HTTPResponse "Successfully created the category"
// @Failure 400 {object} helper.HTTPResponse "Invalid input"
// @Failure 500 {object} helper.HTTPResponse "Failed to create category"
// @Router /category/create [post]
func (ctrl *CategoryController) CreateCatergoryController(c *gin.Context) {
	var categoryInput models.Category
	file, err := c.FormFile("icon")
	if err != nil {
		ctrl.Log.Error("Failed to get uploaded file", zap.Error(err))
		helper.ResponseError(c, err.Error(), "File upload required", http.StatusBadRequest)
		return
	}

	filePath := "./temp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		ctrl.Log.Error("Failed to save uploaded file", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to process file", http.StatusInternalServerError)
		return
	}

	iconURL, err := utils.UploadToCDN(filePath)
	if err != nil {
		ctrl.Log.Error("Failed to upload file to CDN", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to upload to CDN", http.StatusInternalServerError)
		return
	}

	categoryInput.Icon = iconURL

	if err := ctrl.Service.Category.CreateCatergory(categoryInput); err != nil {
		ctrl.Log.Error("Failed to create category", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to create category", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, categoryInput, "Category created successfully", http.StatusCreated)
}

// GetAllCategoriesController godoc
// @Summary Get all categories
// @Description Retrieve all categories with pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} helper.HTTPResponse "List of categories"
// @Router /category/list [get]
func (ctrl *CategoryController) GetAllCategoriesController(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	limitInt64 := int64(limit)

	categories, totalItems, err := ctrl.Service.Category.GetAllCategories(page, limit)
	if err != nil {
		ctrl.Log.Error("Failed to fetch categories", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	totalPages := (totalItems + limitInt64 - 1) / limitInt64
	helper.ResponseOKPagination(c, categories, "Categories fetched successfully", page, limit, int(totalItems), int(totalPages), http.StatusOK)
}

// GetCategoryByIDController godoc
// @Summary Get category by ID
// @Description Retrieve a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category "Category data"
// @Failure 400 {object} helper.HTTPResponse "Invalid category ID"
// @Failure 404 {object} helper.HTTPResponse "Category not found"
// @Router /category/:id [get]
func (ctrl *CategoryController) GetCategoryByIDController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Log.Error("Invalid category ID", zap.Error(err))
		helper.ResponseError(c, "Invalid ID format", "Invalid input", http.StatusBadRequest)
		return
	}

	category, err := ctrl.Service.Category.GetCategoryByID(id)
	if err != nil {
		ctrl.Log.Error("Failed to fetch category by ID", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Category not found", http.StatusNotFound)
		return
	}

	helper.ResponseOK(c, category, "Category fetched successfully", http.StatusOK)
}

// UpdateCategoryController godoc
// @Summary Update a category
// @Description Update the details of an existing category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body models.Category true "Updated category details"
// @Success 200 {object} helper.HTTPResponse "Successfully updated the category"
// @Failure 400 {object} helper.HTTPResponse "Invalid category ID or input"
// @Failure 404 {object} helper.HTTPResponse "Category not found"
// @Failure 500 {object} helper.HTTPResponse "Failed to update category"
// @Router /category/:id [put]
func (ctrl *CategoryController) UpdateCategoryController(c *gin.Context) {
	var categoryInput models.Category
	if err := c.ShouldBind(&categoryInput); err != nil {
		ctrl.Log.Error("Failed to bind input", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid input", http.StatusBadRequest)
		return
	}

	file, err := c.FormFile("icon")
	if err == nil {
		filePath := "./temp/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			ctrl.Log.Error("Failed to save uploaded file", zap.Error(err))
			helper.ResponseError(c, err.Error(), "Failed to process file", http.StatusInternalServerError)
			return
		}

		iconURL, err := utils.UploadToCDN(filePath)
		if err != nil {
			ctrl.Log.Error("Failed to upload file to CDN", zap.Error(err))
			helper.ResponseError(c, err.Error(), "Failed to upload to CDN", http.StatusInternalServerError)
			return
		}

		categoryInput.Icon = iconURL
	}

	if err := ctrl.Service.Category.UpdateCategory(categoryInput); err != nil {
		ctrl.Log.Error("Failed to update category", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to update category", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, categoryInput, "Category updated successfully", http.StatusOK)
}

// DeleteCategoryController godoc
// @Summary Delete category by ID
// @Description Delete a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} helper.HTTPResponse "Successfully deleted the category"
// @Failure 400 {object} helper.HTTPResponse "Invalid category ID"
// @Failure 500 {object} helper.HTTPResponse "Failed to delete category"
// @Router /category/:id [delete]
func (ctrl *CategoryController) DeleteCategoryController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Log.Error("Invalid category ID", zap.Error(err))
		helper.ResponseError(c, "Invalid ID format", "Invalid input", http.StatusBadRequest)
		return
	}

	if err := ctrl.Service.Category.DeleteCategory(id); err != nil {
		ctrl.Log.Error("Failed to delete category", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to delete category", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, nil, "Category deleted successfully", http.StatusOK)
}
