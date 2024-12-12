package bannerservice_test

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/infra"
	"dashboard-ecommerce-team2/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBannerService(t *testing.T) {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		t.Fatalf("Database Error %s\n", err.Error())
	}
	// ctx := infra.MockTest

	t.Run("Create Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Test Banner",
			Type:        models.JSONB{"promo"},
			PathPage:    "/test-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
		}

		err := ctx.Ctl.Banner.Service.Banner.CreateBanner(banner)
		assert.Nil(t, err, "Expected no error during banner creation")
	})

	t.Run("Create Banner - Failure", func(t *testing.T) {
		banner := &models.Banner{} // Missing required fields
		err := ctx.Ctl.Banner.Service.Banner.CreateBanner(banner)
		assert.NotNil(t, err, "Expected error due to missing fields")
		assert.EqualError(t, err, "invalid image url")
	})

	t.Run("Get Banner by ID - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Banner by ID",
			Type:        models.JSONB{"promo", "sale"},
			PathPage:    "/promo-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
		}
		ctx.Ctl.Banner.Service.Banner.CreateBanner(banner)

		result, err := ctx.Ctl.Banner.Service.Banner.GetBannerByID(banner.ID)
		assert.Nil(t, err, "Expected no error while retrieving banner")
		assert.Equal(t, banner.Title, result.Title)
	})

	t.Run("Get Banner by ID - Not Found", func(t *testing.T) {
		result, err := ctx.Ctl.Banner.Service.Banner.GetBannerByID(9999)
		assert.Nil(t, result, "Expected no banner found")
		assert.NotNil(t, err, "Expected error for non-existent banner ID")
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("Update Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:       "/url/example/image.jpg",
			Title:       "Update Test Banner",
			Type:        models.JSONB{"promo"},
			PathPage:    "/update-page",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			Published:   false,
		}
		ctx.Ctl.Banner.Service.Banner.CreateBanner(banner)

		banner.Title = "Updated Title"
		err := ctx.Ctl.Banner.Service.Banner.UpdateBanner(banner)
		assert.Nil(t, err, "Expected no error during banner update")

		updatedBanner, _ := ctx.Ctl.Banner.Service.Banner.GetBannerByID(banner.ID)
		assert.Equal(t, banner.Published, updatedBanner.Published)
	})

	t.Run("Delete Banner - Success", func(t *testing.T) {
		banner := &models.Banner{
			Image:     "/url/example/image.jpg",
			Title:     "Delete Test Banner",
			Type:      models.JSONB{"promo"},
			PathPage:  "/delete-page",
			EndDate:   helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published: true,
		}
		ctx.Ctl.Banner.Service.Banner.CreateBanner(banner)

		err := ctx.Ctl.Banner.Service.Banner.DeleteBanner(banner.ID)
		assert.Nil(t, err, "Expected no error during banner deletion")

		_, err = ctx.Ctl.Banner.Service.Banner.GetBannerByID(banner.ID)
		assert.NotNil(t, err, "Expected error for deleted banner ID")
		assert.Contains(t, err.Error(), "not found")
	})
}
