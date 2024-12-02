package controller

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		service: service.NewProductService(db),
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.service.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
