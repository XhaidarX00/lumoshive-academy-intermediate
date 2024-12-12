package ordercontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewOrderController(service service.Service, log *zap.Logger) *OrderController {
	return &OrderController{
		Service: service,
		Log:     log,
	}
}

// UpdateOrderStatusController godoc
// @Summary Update order status
// @Description Update the status of an order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body string true "Order Status" example("shipped")
// @Success 200 {object} helper.HTTPResponse "Successfully updated the order status"
// @Failure 400 {object} helper.HTTPResponse "Invalid order ID"
// @Failure 500 {object} helper.HTTPResponse "Failed to update order status"
// @Router /orders/{id} [put]
func (ctrl *OrderController) UpdateOrderStatusController(c *gin.Context) {
	idStr := c.Param("id")
	ctrl.Log.Info("Received ID from URL", zap.String("id", idStr))

	if idStr == "" {
		ctrl.Log.Error("Invalid order ID: ID cannot be empty")
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.Log.Error("Invalid order ID format", zap.String("id", idStr), zap.Error(err))
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	var statusRequest struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.BindJSON(&statusRequest); err != nil {
		ctrl.Log.Error("Invalid input", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid input", 400)
		return
	}

	if err := ctrl.Service.Order.UpdateOrderStatus(id, statusRequest.Status); err != nil {
		ctrl.Log.Error("Failed to update order status", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to update order status", 500)
		return
	}

	ctrl.Log.Info("Successfully updated order status", zap.Int("id", id), zap.String("status", statusRequest.Status))
	helper.ResponseOK(c, nil, "Order status updated successfully", 200)
}

// GetAllOrdersController godoc
// @Summary Get all orders with pagination
// @Description Get a paginated list of orders
// @Tags orders
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} helper.HTTPResponse
// @Router /orders/ [get]
func (ctrl *OrderController) GetAllOrdersController(c *gin.Context) {
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	orders, totalItems, err := ctrl.Service.Order.GetAllOrders(page, limit)
	if err != nil {
		ctrl.Log.Error("Failed to fetch paginated orders", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	helper.ResponseOKPagination(c, orders, "Fetched orders successfully", page, limit, int(totalItems), totalPages, http.StatusOK)
}

// GetOrderByIDController godoc
// @Summary Get order by ID
// @Description Get a single order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func (ctrl *OrderController) GetOrderByIDController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Log.Error("Invalid order ID", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := ctrl.Service.Order.GetOrderByID(id)
	if err != nil {
		ctrl.Log.Error("Failed to fetch order by ID", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Order not found", http.StatusNotFound)
		return
	}

	helper.ResponseOK(c, order, "Fetched order by ID successfully", http.StatusOK)
}

// DeleteOrderController godoc
// @Summary Delete order by ID
// @Description Delete an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} helper.HTTPResponse "Successfully deleted the order"
// @Failure 400 {object} helper.HTTPResponse "Invalid order ID"
// @Failure 500 {object} helper.HTTPResponse "Failed to delete the order"
// @Router /orders/{id} [delete]
func (ctrl *OrderController) DeleteOrderController(c *gin.Context) {
	idStr := c.Param("id")
	ctrl.Log.Info("Received ID from URL", zap.String("id", idStr))

	if idStr == "" {
		ctrl.Log.Error("Invalid order ID: ID cannot be empty")
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.Log.Error("Invalid order ID format", zap.String("id", idStr), zap.Error(err))
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	err = ctrl.Service.Order.DeleteOrder(id)
	if err != nil {
		ctrl.Log.Error("Failed to delete order", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to delete order", 500)
		return
	}

	ctrl.Log.Info("Successfully deleted order", zap.Int("id", id))
	helper.ResponseOK(c, nil, "Order deleted successfully", 200)
}

// GetOrderDetailController godoc
// @Summary Get order detail by ID
// @Description Get the details of an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} helper.HTTPResponse
// @Failure 400 {object} helper.HTTPResponse "Invalid order ID"
// @Failure 500 {object} helper.HTTPResponse "Failed to fetch order details"
// @Router /orders/detail/{id} [get]
func (ctrl *OrderController) GetOrderDetailController(c *gin.Context) {
	idStr := c.Param("id")
	ctrl.Log.Info("Received ID from URL", zap.String("id", idStr))

	if idStr == "" {
		ctrl.Log.Error("Invalid order ID: ID cannot be empty")
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.Log.Error("Invalid order ID format", zap.String("id", idStr), zap.Error(err))
		helper.ResponseError(c, "Invalid order ID", "Invalid order ID", 400)
		return
	}

	order, orderItems, err := ctrl.Service.Order.GetOrderDetail(id)
	if err != nil {
		ctrl.Log.Error("Failed to fetch order details", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to fetch order details", 500)
		return
	}

	helper.ResponseOK(c, models.OrderDetailResponse{Order: order, Items: orderItems}, "Fetched order details successfully", http.StatusOK)
}
