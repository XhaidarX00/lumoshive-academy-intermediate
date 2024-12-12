package productservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	utils "dashboard-ecommerce-team2/util"
	"errors"
	"os"

	"go.uber.org/zap"
)

type ProductService interface {
	CreateProduct(product *models.Product, filePath string) (*models.Product, error)
	GetAllProducts(page, pageSize int) ([]*models.ProductWithCategory, int, error)
	GetProductByID(id int) (*models.ProductID, error)
	UpdateProduct(id int, product models.Product, filePath string) (*models.Product, error)
	DeleteProduct(id int) error
}

type productService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(product *models.Product, filePath string) (*models.Product, error) {
	cdnURL, err := utils.UploadToCDN(filePath)
	if err != nil {
		p.Log.Error("service: upload failed", zap.Error(err))
		return product, err
	}

	product.Images = []string{cdnURL}

	if err := p.Repo.Product.Create(product); err != nil {
		p.Log.Error("service: create failed", zap.Error(err))
		return product, err
	}

	os.Remove(filePath)

	return product, nil
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(id int) error {
	err := p.Repo.Product.Delete(id)
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

// GetAllProducts implements ProductService.
func (p *productService) GetAllProducts(page, pageSize int) ([]*models.ProductWithCategory, int, error) {
	products, totalItems, err := p.Repo.Product.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	if len(products) == 0 {
		return nil, int(totalItems), errors.New("product not found")
	}
	// Convert totalItems from int64 to int before returning
	return products, int(totalItems), nil
}

// GetProductByID implements ProductService.
func (p *productService) GetProductByID(id int) (*models.ProductID, error) {
	product, err := p.Repo.Product.GetByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// UpdateProduct implements ProductService.
func (p *productService) UpdateProduct(id int, product models.Product, filePath string) (*models.Product, error) {
	if filePath != "" {
        cdnURL, err := utils.UploadToCDN(filePath)
        if err != nil {
            p.Log.Error("service: upload failed", zap.Error(err))
            return nil, err
        }
        product.Images = []string{cdnURL}
    }

    updatedProduct, err := p.Repo.Product.Update(id, product)
    if err != nil {
        p.Log.Error("service: update failed", zap.Error(err))
        return nil, err
    }

    os.Remove(filePath) 
    return updatedProduct, nil
}

func NewProductService(repo repository.Repository, log *zap.Logger) ProductService {
	return &productService{Repo: repo, Log: log}
}
