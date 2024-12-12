package promotioncontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PromotionController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewPromotionController(service service.Service, log *zap.Logger) *PromotionController {
	return &PromotionController{
		Service: service,
		Log:     log,
	}
}

// GetAllPromotionsController godoc
// @Summary Get all promotions
// @Description Retrieve a list of all promotions
// @Tags Promotions
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.Promotion} "Successfully retrieved promotions"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/list-promotion [get]
func (ctrl *PromotionController) GetAllPromotionsController(c *gin.Context) {
	promotion, err := ctrl.Service.Promotion.GetAllPromotions()
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, promotion, "Successfully get promotion", http.StatusOK)
}

// GetByIdPromotionsController godoc
// @Summary Get promotion by ID
// @Description Retrieve a specific promotion by its ID
// @Tags Promotions
// @Produce json
// @Param id query int true "Promotion ID"
// @Success 200 {object} models.SuccessResponse{data=[]models.Promotion} "Successfully retrieved promotion"
// @Failure 400 {object} models.ErrorResponse "Invalid promotion ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/promotion [get]
func (ctrl *PromotionController) GetByIdPromotionsController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id promotion invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	promotion, err := ctrl.Service.Promotion.GetByIDPromotion(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, promotion, "Success get promotion", http.StatusOK)
}

// CreatePromotionController godoc
// @Summary Create a new promotion
// @Description Create a new promotion with detailed information
// @Tags Promotions
// @Accept json
// @Produce json
// @Param promotion body models.SuccessResponse{data=[]models.Promotion} true "Promotion Details"
// @Success 200 {object} models.SuccessResponse{data=[]models.Promotion} "Successfully created promotion"
// @Failure 400 {object} models.ErrorResponse "Invalid input data"
// @Failure 500 {object} models.ErrorResponse "Failed to create promotion"
// @Router /api/promotions [post]
func (ctrl *PromotionController) CreatePromotionController(c *gin.Context) {
	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	if err := ctrl.Service.Promotion.CreatePromotion(&promotion); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create promotion", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, promotion, "Successfully created promotion", http.StatusOK)
}

// UpdatePromotionController godoc
// @Summary Update promotion published status
// @Description Toggle the published status of a specific promotion
// @Tags Promotions
// @Produce json
// @Param id query int true "Promotion ID"
// @Success 200 {object} models.SuccessResponse{data=[]models.Promotion} "Successfully updated promotion status"
// @Failure 400 {object} models.ErrorResponse "Invalid promotion ID"
// @Failure 500 {object} models.ErrorResponse "Failed to update promotion"
// @Router /api/promotions/status [put]
func (ctrl *PromotionController) UpdatePromotionController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "promotion id invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	promotion, err := ctrl.Service.Promotion.GetByIDPromotion(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	// Toggle published status
	promotion.Published = !promotion.Published
	err = ctrl.Service.Promotion.UpdatePromotion(promotion)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, promotion, "Successfully update published promotion", http.StatusOK)
}

// DeletePromotionController godoc
// @Summary Delete a promotion
// @Description Delete a specific promotion by its ID
// @Tags Promotions
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Role"
// @Param id query int true "Promotion ID"
// @Success 200 {object} models.SuccessResponse "Successfully deleted promotion"
// @Failure 400 {object} models.ErrorResponse "Invalid promotion ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Forbidden - Insufficient privileges"
// @Failure 500 {object} models.ErrorResponse "Failed to delete promotion"
// @Router /api/promotions [delete]
func (ctrl *PromotionController) DeletePromotionController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id promotion invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	err = ctrl.Service.Promotion.DeletePromotion(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "Successfully delete promotion", http.StatusOK)
}
