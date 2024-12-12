package productrepository

import (
	"dashboard-ecommerce-team2/models"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(productInput *models.Product) error
	Update(id int, productInput models.Product) (*models.Product, error)
	Delete(id int) error
	GetByID(id int) (*models.ProductID, error)
	GetAll(page, pageSize int) ([]*models.ProductWithCategory, int64, error)
	CountProduct() (int, error)
	CountEachProduct() ([]models.BestProduct, error)
}

type productRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewProductRepository(db *gorm.DB, log *zap.Logger) ProductRepository {
	return &productRepository{DB: db, Log: log}
}

// CountEachProduct implements ProductRepository.
func (p *productRepository) CountEachProduct() ([]models.BestProduct, error) {
	type ProductCount struct {
		ProductID uint   `gorm:"column:product_id"`
		Name      string `gorm:"column:name"`
		Total     int    `gorm:"column:total"`
	}

	var productCounts []ProductCount
	var bestProducts []models.BestProduct

	// Perform the query to count products grouped by product ID
	err := p.DB.Table("order_items").
		Select("order_items.product_id, products.name, COALESCE(COUNT(order_items.product_id), 0) AS total").
		Joins("JOIN orders ON orders.id = order_items.order_id AND orders.status = ?", "completed").
		Joins("JOIN products ON products.id = order_items.product_id").
		Group("order_items.product_id, products.name").
		Order("total DESC").
		Scan(&productCounts).Error

	if err != nil {
		return nil, err
	}

	// Map the query results to the []models.BestProduct slice
	for _, pc := range productCounts {
		bestProducts = append(bestProducts, models.BestProduct{
			ProductID: int(pc.ProductID),
			Name:      pc.Name,
			Total:     pc.Total,
		})
	}

	return bestProducts, nil
}

// CountProduct implements ProductRepository.
func (p *productRepository) CountProduct() (int, error) {
	var count int64

	// Query to count distinct products in completed orders
	err := p.DB.Table("order_items").
		Joins("JOIN orders ON orders.id = order_items.order_id AND orders.status = ?", "completed").
		Where("EXTRACT(MONTH FROM orders.created_at) = ? AND EXTRACT(YEAR FROM orders.created_at) = ?", time.Now().Month(), time.Now().Year()).
		Select("COALESCE(SUM(order_items.quantity),0)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// Create implements ProductRepository.
func (p *productRepository) Create(productInput *models.Product) error {

	p.Log.Info("Creating product", zap.Any("input", productInput))
	err := p.DB.Create(productInput).Error
	if err != nil {
		p.Log.Error("repository: Error creating product", zap.Error(err))
		return err
	}
	return nil
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id int) error {
	var product models.Product

	err := p.DB.First(&product, id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("product not found")
		}
		p.Log.Error("Failed to fetch product for deletion", zap.Error(err))
		return err
	}

	err = p.DB.Delete(&product).Error
	if err != nil {
		p.Log.Error("Failed to delete product", zap.Error(err))
		return err
	}

	return nil
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll(page, pageSize int) ([]*models.ProductWithCategory, int64, error) {
	// var products []*models.Product
	var productsWithCategory []*models.ProductWithCategory
	var totalItems int64

	offset := (page - 1) * pageSize

	// Menghitung total items
	err := p.DB.Model(&models.Product{}).Count(&totalItems).Error
	if err != nil {
		p.Log.Error("Failed to count total products", zap.Error(err))
		return nil, 0, err
	}

	// Mengambil produk dengan pagination
	err = p.DB.Table("products").
		Select(`products.id, 
            categories.name AS category_name, 
            products.name, 
            products.code_product, 
            products.images, 
            products.description, 
            products.stock, 
            products.price, 
            products.published`).
		Joins("JOIN categories ON categories.id = products.category_id").
		Offset(offset).
		Limit(pageSize).
		Find(&productsWithCategory).Error

	if err != nil {
		p.Log.Error("Failed to fetch products", zap.Error(err))
		return nil, 0, err
	}

	return productsWithCategory, totalItems, nil
}

// GetByID implements ProductRepository.
func (p *productRepository) GetByID(id int) (*models.ProductID, error) {
	var productID models.ProductID
	var variantJSON []byte

	err := p.DB.Table("products").
		Select(`
			products.id, 
			products.category_id,
			categories.name AS category_name, 
			CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant,
			products.name, 
			products.code_product, 
			products.images, 
			products.description, 
			products.stock, 
			products.price, 
			products.published
		`).
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.id = ?", id).
		First(&productID).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("product not found")
		}
		p.Log.Error("Failed to fetch product by ID", zap.Error(err))
		return nil, err
	}
	if len(variantJSON) > 0 {
		if err := json.Unmarshal(variantJSON, &productID.Variant); err != nil {
			p.Log.Error("Repository: failed to unmarshal variant JSON", zap.Error(err))
			return nil, err
		}
	}

	return &productID, nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(id int, productInput models.Product) (*models.Product, error) {
	p.Log.Info("Updated product", zap.Any("input", productInput))

	var product models.Product
	err := p.DB.Model(&models.Product{}).Where("id = ?", id).Updates(productInput).First(&product).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("product not found")
		}
		p.Log.Error("repository: update failed", zap.Error(err))
		return nil, err
	}
	return &product, nil
}
