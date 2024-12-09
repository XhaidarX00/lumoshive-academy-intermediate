package banners

import (
	"net/http"
	"strconv"
	"voucher_system/helper"
	"voucher_system/models"
	"voucher_system/service/banner"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	Svc banner.BannerServiceInterface
}

func NewBannerHandler(svc banner.BannerServiceInterface) *BannerHandler {
	return &BannerHandler{Svc: svc}
}

func (h *BannerHandler) CreateBanner(c *gin.Context) {
	var banner models.Banner

	// Menggunakan ShouldBind untuk mendukung berbagai jenis input (JSON, form-urlencoded, atau multipart)
	if err := c.ShouldBind(&banner); err != nil {
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	if err := h.Svc.Create(&banner); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create banner", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, banner, "succes created banner", http.StatusOK)
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid banner ID"})
		return
	}

	var banner models.Banner

	// Menggunakan ShouldBind agar mendukung input form
	if err := c.ShouldBind(&banner); err != nil {
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	banner.ID = uint(id)
	if err := h.Svc.Update(&banner); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create banner", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "succers update banner", http.StatusOK)
}
