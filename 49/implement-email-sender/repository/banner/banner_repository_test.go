package bannerrepository_test

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/infra"
	"dashboard-ecommerce-team2/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBannerRepository(t *testing.T) {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		t.Fatalf("Database Error %s\n", err.Error())
	}

	// ctx := infra.MockTest
	repo := ctx.Repo.Banner

	// Create
	t.Run("Create Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Test Banner",
			Type:        models.JSONB{"seasonal", "promo"},
			PathPage:    "/spring-sale",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
		}

		err := repo.Create(banner)
		assert.Nil(t, err, "Expected no error during banner creation")
	})

	t.Run("Create Banner - Failure", func(t *testing.T) {
		banner := &models.Banner{} // Missing required fields
		err := repo.Create(banner)
		assert.EqualError(t, err, "invalid image url")
	})

	t.Run("Create Banner - Failure unix constraint", func(t *testing.T) {
		banner := &models.Banner{ID: 1} // Missing required fields
		err := repo.Create(banner)
		assert.EqualError(t, err, "invalid image url")
	})

	// GetByID
	t.Run("Get Banner by ID - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Banner by ID",
			Type:        models.JSONB{"promo"},
			PathPage:    "/promo-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
		}
		repo.Create(banner)

		result, err := repo.GetByID(banner.ID)
		assert.Nil(t, err, "Expected no error while retrieving banner")
		assert.Equal(t, banner.Title, result.Title)
	})

	t.Run("Get Banner by ID - Not Found", func(t *testing.T) {
		result, err := repo.GetByID(9999)
		assert.Nil(t, result, "Expected no banner found")
		assert.NotNil(t, err, "Expected error for non-existent banner ID")
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("Get Banner by ID - Not Found", func(t *testing.T) {
		result, err := repo.GetByID(1)
		assert.Nil(t, result, "Expected no banner found")
		assert.NotNil(t, err, "Expected error for non-existent banner ID")
		assert.Contains(t, err.Error(), "not found")
	})

	// Update
	t.Run("Update Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Update Test Banner",
			Type:        models.JSONB{"promo"},
			PathPage:    "/update-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   false,
		}
		repo.Create(banner)

		banner.Title = "Updated Title"
		err := repo.Update(banner)
		assert.Nil(t, err, "Expected no error during banner update")

		updatedBanner, _ := repo.GetByID(banner.ID)
		assert.Equal(t, banner.Published, updatedBanner.Published)
	})

	t.Run("Update Banner - Failure", func(t *testing.T) {
		banner := &models.Banner{ID: 1}
		err := repo.Update(banner)
		assert.EqualError(t, err, "banner with ID 1 not found")
	})

	// Delete
	t.Run("Delete Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Delete Test Banner",
			Type:        models.JSONB{"promo"},
			PathPage:    "/delete-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
		}
		repo.Create(banner)

		err := repo.Delete(banner.ID)
		assert.Nil(t, err, "Expected no error during banner deletion")

		_, err = repo.GetByID(banner.ID)
		assert.NotNil(t, err, "Expected error for deleted banner ID")
		assert.Contains(t, err.Error(), "not found")
	})
	t.Run("Delete Banner - Failure Id Not Found", func(t *testing.T) {
		err := repo.Delete(1)
		assert.EqualError(t, err, "banner with ID 1 not found")
	})

	t.Run("Delete Banner - Failure Id Not Found", func(t *testing.T) {
		err := repo.Delete(99999)
		assert.EqualError(t, err, "banner with ID 99999 not found")
	})
}
