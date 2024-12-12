package categoryrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(categoryInput models.Category) error
	Update(categoryInput models.Category) error
	Delete(id int) error
	GetByID(id int) (*models.Category, error)
	GetAll(page, limit int) ([]models.Category, int64, error)
}

type categoryRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements CategoryRepository.
func (c *categoryRepository) Create(categoryInput models.Category) error {
	// var existingCategory models.Category
	// if err := c.DB.Where("name = ?", categoryInput.Name).First(&existingCategory).Error; err == nil {
	// 	log.Printf("Category with name '%s' already exists", categoryInput.Name)
	// 	return errors.New("category already exists")
	// }

	c.Log.Info("Creating a new category", zap.String("name", categoryInput.Name), zap.String("icon", categoryInput.Icon))
	if err := c.DB.Create(&categoryInput).Error; err != nil {
		c.Log.Error("Failed to create category", zap.Error(err))
		return err
	}
	return nil
}

// Delete implements CategoryRepository.
func (c *categoryRepository) Delete(id int) error {
	c.Log.Info("Deleting category", zap.Int("id", id))
	if err := c.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.Log.Error("Failed to delete category", zap.Error(err))
		return err
	}
	return nil
}

// GetAll implements CategoryRepository.
func (c *categoryRepository) GetAll(page, limit int) ([]models.Category, int64, error) {
	c.Log.Info("Fetching all categories with pagination", zap.Int("page", page), zap.Int("limit", limit))

	var categories []models.Category
	var totalItems int64

	if err := c.DB.Model(&models.Category{}).Count(&totalItems).Error; err != nil {
		c.Log.Error("Failed to count categories", zap.Error(err))
		return nil, 0, err
	}

	offset := (page - 1) * limit

	if err := c.DB.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		c.Log.Error("Failed to fetch paginated categories", zap.Error(err))
		return nil, 0, err
	}

	return categories, totalItems, nil
}

// GetByID implements CategoryRepository.
func (c *categoryRepository) GetByID(id int) (*models.Category, error) {
	c.Log.Info("Fetching category by ID", zap.Int("id", id))
	var category models.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		c.Log.Error("Failed to fetch category by ID", zap.Error(err))
		return nil, err
	}
	return &category, nil
}

// Update implements CategoryRepository.
func (c *categoryRepository) Update(categoryInput models.Category) error {
	c.Log.Info("Updating category", zap.Int("id", categoryInput.ID))
	if err := c.DB.Model(&models.Category{}).Where("id = ?", categoryInput.ID).Updates(categoryInput).Error; err != nil {
		c.Log.Error("Failed to update category", zap.Error(err))
		return err
	}
	return nil
}

func NewCategoryRepository(db *gorm.DB, log *zap.Logger) CategoryRepository {
	return &categoryRepository{DB: db, Log: log}
}
