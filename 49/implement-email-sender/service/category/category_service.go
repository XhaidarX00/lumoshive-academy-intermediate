package categoryservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type CategoryService interface {
	CreateCatergory(categoryInput models.Category) error
	GetAllCategories(page, limit int) ([]models.Category, int64, error)
	GetCategoryByID(id int) (*models.Category, error)
	UpdateCategory(categoryInput models.Category) error
	DeleteCategory(id int) error
}

type categoryService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreateCatergory implements CategoryService.
func (c *categoryService) CreateCatergory(categoryInput models.Category) error {
	c.Log.Info("Service: Creating category", zap.String("name", categoryInput.Name))
	if err := c.Repo.Category.Create(categoryInput); err != nil {
		c.Log.Error("Service: Failed to create category", zap.Error(err))
		return err
	}
	return nil
}

// DeleteCategory implements CategoryService.
func (c *categoryService) DeleteCategory(id int) error {
	c.Log.Info("Service: Deleting category", zap.Int("id", id))
	if err := c.Repo.Category.Delete(id); err != nil {
		c.Log.Error("Service: Failed to delete category", zap.Error(err))
		return err
	}
	return nil
}

// GetAllCategories implements CategoryService.
func (c *categoryService) GetAllCategories(page, limit int) ([]models.Category, int64, error) {
	c.Log.Info("Service: Fetching all categories", zap.Int("page", page), zap.Int("limit", limit))
	categories, totalItems, err := c.Repo.Category.GetAll(page, limit)
	if err != nil {
		c.Log.Error("Service: Failed to fetch categories", zap.Error(err))
		return nil, 0, err
	}
	return categories, totalItems, nil
}

// GetCategoryByID implements CategoryService.
func (c *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	c.Log.Info("Service: Fetching category by ID", zap.Int("id", id))
	category, err := c.Repo.Category.GetByID(id)
	if err != nil {
		c.Log.Error("Service: Failed to fetch category by ID", zap.Error(err))
		return nil, err
	}
	return category, nil
}

// UpdateCategory implements CategoryService.
func (c *categoryService) UpdateCategory(categoryInput models.Category) error {
	c.Log.Info("Service: Updating category", zap.Int("id", categoryInput.ID))
	if err := c.Repo.Category.Update(categoryInput); err != nil {
		c.Log.Error("Service: Failed to update category", zap.Error(err))
		return err
	}
	return nil
}

func NewCategoryService(repo repository.Repository, log *zap.Logger) CategoryService {
	return &categoryService{Repo: repo, Log: log}
}
