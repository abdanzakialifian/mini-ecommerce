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
	orderService order.Service
}

func NewHandler(orderService order.Service) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Create(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	var newItems []order.NewItem
	for _, item := range req.Items {
		newItem := order.NewItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		newItems = append(newItems, newItem)
	}

	orderDetail, appErr := h.orderService.Create(c.Request.Context(), userId, newItems)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	itemResponses := []ItemResponse{}
	for _, item := range orderDetail.Items {
		itemResponse := ItemResponse{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Price:     item.Price,
			Quantity:  item.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	status, res := response.Success(
		"Success Create Order",
		DetailResponse{
			Order: Response{
				ID:         orderDetail.Data.ID,
				UserID:     orderDetail.Data.UserID,
				TotalPrice: orderDetail.Data.TotalPrice,
				Status:     orderDetail.Data.Status,
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

	orderDetail, appErr := h.orderService.Get(c.Request.Context(), orderId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	itemResponses := []ItemResponse{}
	for _, item := range orderDetail.Items {
		itemResponse := ItemResponse{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Price:     item.Price,
			Quantity:  item.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	status, res := response.Success(
		"Success Get Order",
		DetailResponse{
			Order: Response{
				ID:         orderDetail.Data.ID,
				UserID:     orderDetail.Data.UserID,
				TotalPrice: orderDetail.Data.TotalPrice,
				Status:     orderDetail.Data.Status,
			},
			Items: itemResponses,
		},
	)
	c.JSON(status, res)
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	orderDetails, appErr := h.orderService.GetByUserId(c.Request.Context(), userId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var detailResponses []DetailResponse
	for _, orderDetail := range orderDetails {
		itemResponses := []ItemResponse{}
		for _, item := range orderDetail.Items {
			itemResponse := ItemResponse{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ProductID: item.ProductID,
				Price:     item.Price,
				Quantity:  item.Quantity,
			}
			itemResponses = append(itemResponses, itemResponse)
		}

		detailResponse := DetailResponse{
			Order: Response{
				ID:         orderDetail.Data.ID,
				UserID:     orderDetail.Data.UserID,
				TotalPrice: orderDetail.Data.TotalPrice,
				Status:     orderDetail.Data.Status,
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

	if appErr := h.orderService.UpdateStatus(c.Request.Context(), orderId, req.Status); appErr != nil {
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

	if appErr := h.orderService.Cancel(c.Request.Context(), orderId); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Cancelled Order")
	c.JSON(status, res)
}
