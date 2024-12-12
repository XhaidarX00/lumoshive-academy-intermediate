package bannercontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewBannerController(service service.Service, log *zap.Logger) *BannerController {
	return &BannerController{
		Service: service,
		Log:     log,
	}
}

// CreateBannerController godoc
// @Summary Create a new banner
// @Description Create a new banner with image upload
// @Tags Banner
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Banner Title"
// @Param description formData string false "Banner Description"
// @Param image_path formData file true "Banner Image"
// @Success 200 {object} models.SuccessResponse{data=models.Banner} "Successfully created banner"
// @Failure 400 {object} models.ErrorResponse "Invalid data input"
// @Failure 500 {object} models.ErrorResponse "Failed to upload image or create banner"
// @Router /api/create-banner [post]
func (ctrl *BannerController) CreateBannerController(c *gin.Context) {
	var banner models.BannerGetValue

	// Bind the form data to the banner struct (without the file)
	if err := c.ShouldBind(&banner); err != nil {
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	// Process image upload
	imageURL, err := helper.UploadImage(banner.ImagePath)
	if err != nil {
		log.Println(err.Error())
		helper.ResponseError(c, "FAILED", "Failed to upload image", http.StatusInternalServerError)
		return
	}

	// Assign the image URL to the banner
	var bannerSave models.Banner
	bannerSave, err = bannerSave.CopyBannerGetValueToBanner(imageURL, banner)
	if err != nil {
		log.Println("Error binding form data:", err.Error())
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	// Call service to create the banner
	if err := ctrl.Service.Banner.CreateBanner(&bannerSave); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create banner", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, bannerSave, "Successfully created banner", http.StatusOK)
}

// UpdateBannerController godoc
// @Summary Update banner published status
// @Description Toggle the published status of a banner
// @Tags Banner
// @Accept multipart/form-data
// @Produce json
// @Param id formData int true "Banner ID"
// @Success 200 {object} models.SuccessResponse{data=models.Banner} "Successfully updated published banner"
// @Failure 400 {object} models.ErrorResponse "Invalid banner ID"
// @Failure 500 {object} models.ErrorResponse "Failed to update banner"
// @Router /api/banner [put]
func (ctrl *BannerController) UpdateBannerController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	// banner, err := ctrl.Service.Banner.GetBannerByID(idInt)
	// if err != nil {
	// 	helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
	// 	c.Abort()
	// 	return
	// }

	// change status publish
	// banner.Published = !banner.Published

	var banner models.Banner
	banner.ID = idInt
	err = ctrl.Service.Banner.UpdateBanner(&banner)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, banner, "Successfully update published banner", http.StatusOK)
}

// GetBannerByIDController godoc
// @Summary Get banner by ID
// @Description Retrieve a specific banner by its ID
// @Tags Banner
// @Accept multipart/form-data
// @Produce json
// @Param id formData int true "Banner ID"
// @Success 200 {object} models.SuccessResponse{data=models.Banner} "Successfully retrieved banner"
// @Failure 400 {object} models.ErrorResponse "Invalid banner ID"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve banner"
// @Router /api/banner [get]
func (ctrl *BannerController) GetBannerByIDController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	banner, err := ctrl.Service.Banner.GetBannerByID(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, banner, "Successfully get banner", http.StatusOK)
}

// DeleteBannerController godoc
// @Summary Delete a banner
// @Description Remove a banner by its ID
// @Tags Banner
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Role"
// @Param id formData int true "Banner ID"
// @Success 200 {object} models.SuccessResponse "Successfully deleted banner"
// @Failure 400 {object} models.ErrorResponse "Invalid banner ID"
// @Failure 500 {object} models.ErrorResponse "Failed to delete banner"
// @Router /api/banner [delete]
func (ctrl *BannerController) DeleteBannerController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	err = ctrl.Service.Banner.DeleteBanner(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "Successfully delete banner", http.StatusOK)
}
