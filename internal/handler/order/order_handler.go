package order

import (
	"errors"
	"mini-ecommerce/internal/domain/order"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService order.OrderService
}

func NewHandler(orderService order.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Create(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req CreateRequest
	if len(req.Items) == 0 {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Order items cannot be empty",
			nil,
		))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	var items []order.CreateOrderItem
	for _, item := range req.Items {
		orderItem := order.CreateOrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}

		items = append(items, orderItem)
	}

	orderDetail, appErr := h.orderService.CreateOrder(c.Request.Context(), userId, items)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	itemResponses := []ItemResponse{}
	for _, detail := range orderDetail.Items {
		itemResponse := ItemResponse{
			ID:        detail.ID,
			OrderID:   detail.OrderID,
			ProductID: detail.ProductID,
			Price:     detail.Price,
			Quantity:  detail.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	status, res := response.Success(
		"Success Create Order",
		DetailResponse{
			Order: Response{
				ID:         orderDetail.Order.ID,
				UserID:     orderDetail.Order.UserID,
				TotalPrice: orderDetail.Order.TotalPrice,
				Status:     orderDetail.Order.Status,
			},
			Items: itemResponses,
		},
	)
	c.JSON(status, res)
}

func (h *OrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Order id is required"),
		))
		return
	}

	orderId, err := strconv.Atoi(id)

	if err != nil {
		c.Error(helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			errors.New("Error convert oder id"),
		))
		return
	}

	detail, appErr := h.orderService.GetOrder(c.Request.Context(), orderId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	itemResponses := []ItemResponse{}
	for _, detail := range detail.Items {
		itemResponse := ItemResponse{
			ID:        detail.ID,
			OrderID:   detail.OrderID,
			ProductID: detail.ProductID,
			Price:     detail.Price,
			Quantity:  detail.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	status, res := response.Success(
		"Success Get Order",
		DetailResponse{
			Order: Response{
				ID:         detail.Order.ID,
				UserID:     detail.Order.UserID,
				TotalPrice: detail.Order.TotalPrice,
				Status:     detail.Order.Status,
			},
			Items: itemResponses,
		},
	)
	c.JSON(status, res)
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	details, appErr := h.orderService.GetOrderByUserId(c.Request.Context(), userId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var detailResponses []DetailResponse
	for _, detail := range details {
		itemResponses := []ItemResponse{}
		for _, detail := range detail.Items {
			itemResponse := ItemResponse{
				ID:        detail.ID,
				OrderID:   detail.OrderID,
				ProductID: detail.ProductID,
				Price:     detail.Price,
				Quantity:  detail.Quantity,
			}
			itemResponses = append(itemResponses, itemResponse)
		}

		detailResponse := DetailResponse{
			Order: Response{
				ID:         detail.Order.ID,
				UserID:     detail.Order.UserID,
				TotalPrice: detail.Order.TotalPrice,
				Status:     detail.Order.Status,
			},
			Items: itemResponses,
		}

		detailResponses = append(detailResponses, detailResponse)
	}

	status, res := response.Success(
		"Success Get Orders",
		detailResponses,
	)
	c.JSON(status, res)
}

func (h *OrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Order id is required"),
		))
		return
	}

	orderId, err := strconv.Atoi(id)

	if err != nil {
		c.Error(helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			errors.New("Error convert oder id"),
		))
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	appErr := h.orderService.UpdateOrderStatus(c.Request.Context(), orderId, req.Status)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Update Status")
	c.JSON(status, res)
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Order id is required"),
		))
		return
	}

	orderId, err := strconv.Atoi(id)

	if err != nil {
		c.Error(helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			errors.New("Error convert oder id"),
		))
		return
	}

	appErr := h.orderService.CancelOrder(c.Request.Context(), orderId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Cancelled Order")
	c.JSON(status, res)
}
