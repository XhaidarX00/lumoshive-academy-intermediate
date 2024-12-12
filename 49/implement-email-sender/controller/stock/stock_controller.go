package stockcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewStockController(service service.Service, log *zap.Logger) *StockController {
	return &StockController{
		Service: service,
		Log:     log,
	}
}

// UpdateProductStockController godoc
// @Summary      Update product stock
// @Description  Update the stock of a product (increase or decrease).
// @Tags         Stock
// @Accept       json
// @Produce      json
// @Param        stock_request  body      models.StockRequest  true  "Stock update request"
// @Success      200            {object}  utils.ResponseOK "Update success"
// @Failure      400            {object}  utils.ErrorResponse "Invalid input"
// @Failure      500            {object}  utils.ErrorResponse "Internal server error"
// @Security Authentication
// @Security UserID
// @Router       /stock [put]
func (ctrl *StockController) UpdateProductStockController(c *gin.Context) {
	var req models.StockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Log.Warn("invalid input", zap.Error(err))
		helper.ResponseError(c, "Invalid input", helper.FormatValidationError(err), http.StatusBadRequest)
		return
	}

	if req.Type != "in" && req.Type != "out" {
		ctrl.Log.Warn("invalid stock type", zap.String("type", req.Type))
		helper.ResponseError(c, "Invalid stock type", "Type must be 'in' or 'out'", http.StatusBadRequest)
		return
	}

	if err := ctrl.Service.Stock.UpdateProductStock(&req); err != nil {
		ctrl.Log.Error("failed to update stock", zap.Error(err))
		helper.ResponseError(c, "Failed to update stock", err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, nil, "Stock updated successfully", http.StatusOK)
}

// GetProductStockDetailController godoc
// @Summary      Get product stock details
// @Description  Retrieve the stock details of a specific product by stock history ID.
// @Tags         Stock
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Stock history ID"
// @Success      200  {object}  utils.ResponseOK{data=utils.StockResponse} "Detail Stock history"
// @Failure      404  {object}  utils.ErrorResponse "Stock History not found"
// @Security Authentication
// @Security UserID
// @Router       /stock/{id} [get]
func (ctrl *StockController) GetProductStockDetailController(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))
	stockHistory, err := ctrl.Service.Stock.GetProductStockDetail(id)
	if err != nil {
		ctrl.Log.Warn("handler: No history found", zap.Int("id", id))
		helper.ResponseError(c, "No history found", err.Error(), http.StatusNotFound)
		return
	}

	helper.ResponseOK(c, stockHistory, "", http.StatusOK)
}

// DeleteProductStockController godoc
// @Summary      Delete stock history
// @Description  Delete a specific stock history record by its ID.
// @Tags         Stock
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Stock history ID"
// @Success      200  {object}  utils.ResponseOK "Delete Success"
// @Failure      404  {object}  utils.ErrorResponse "Stock history not found"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Security Authentication
// @Security UserID
// @Security UserRole
// @Router       /stock/{id} [delete]
func (ctrl *StockController) DeleteProductStockController(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))

	err := ctrl.Service.Stock.DeleteProductStock(id)
	if err != nil {
		if err.Error() == "history not found" {
			ctrl.Log.Warn("handler: No history found", zap.Int("id", id))
			helper.ResponseError(c, "No history found", err.Error(), http.StatusNotFound)
		} else {
			ctrl.Log.Debug("handler: Failed to delete history", zap.Int("id", id), zap.Error(err))
			ctrl.Log.Error("handler: Failed to delete history", zap.Int("id", id), zap.Error(err))
			helper.ResponseError(c, "Failed", err.Error(), http.StatusInternalServerError)
		}
		return
	}
	helper.ResponseOK(c, nil, "history deleted successfully", http.StatusOK)
}
