package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/order"
	"mini-ecommerce/internal/domain/product"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type orderServiceImpl struct {
	tx                  *helper.Transaction
	orderRepository     order.OrderRepository
	orderItemRepository order.OrderItemRepository
	productRepository   product.Repository
}

func NewOrderService(tx *helper.Transaction, orderRepository order.OrderRepository, orderItemRepository order.OrderItemRepository) order.OrderService {
	return &orderServiceImpl{tx: tx, orderRepository: orderRepository, orderItemRepository: orderItemRepository}
}

func (o *orderServiceImpl) CreateOrder(ctx context.Context, userId int, createOrderItems []order.CreateOrderItem) (order.OrderDetail, *helper.AppError) {
	var orderDetail order.OrderDetail
	err := o.tx.ExecTx(ctx, func(ctx context.Context) error {
		var totalPrice float64
		var results []order.OrderItem

		for _, item := range createOrderItems {
			product, err := o.productRepository.Find(ctx, item.ProductID)
			if err != nil {
				return err
			}

			if product.Stock < item.Quantity {
				return helper.ErrProductInsufficientStock
			}

			itemPrice := product.Price * float64(item.Quantity)
			totalPrice += itemPrice

			orderItem := order.OrderItem{
				ProductID: item.ProductID,
				Price:     product.Price,
				Quantity:  item.Quantity,
			}

			results = append(results, orderItem)
		}

		orderEntity := order.Order{
			UserID:     userId,
			TotalPrice: totalPrice,
			Status:     order.StatusPending,
		}

		if err := o.orderRepository.Create(ctx, &orderEntity); err != nil {
			return err
		}

		for i := range results {
			results[i].OrderID = orderEntity.ID
		}

		if err := o.orderItemRepository.CreateOrderItems(ctx, results); err != nil {
			return err
		}

		for _, orderItem := range results {
			err := o.productRepository.UpdateStock(ctx, orderItem.ProductID, orderItem.Quantity)
			if err != nil {
				return err
			}
		}

		orderDetail = order.OrderDetail{
			Order: order.Order{
				ID:         orderEntity.ID,
				UserID:     orderEntity.UserID,
				TotalPrice: orderEntity.TotalPrice,
				Status:     orderEntity.Status,
			},
			Items: results,
		}
		return nil
	})

	if err != nil {
		if errors.Is(err, helper.ErrProductNotFound) {
			return orderDetail, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		if errors.Is(err, helper.ErrProductInsufficientStock) {
			return orderDetail, helper.NewAppError(
				http.StatusConflict,
				"Insufficient Stock",
				err,
			)
		}

		return orderDetail, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return orderDetail, nil
}

func (o *orderServiceImpl) GetOrder(ctx context.Context, id int) (order.OrderDetail, *helper.AppError) {
	orderEntity, err := o.orderRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return order.OrderDetail{}, helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return order.OrderDetail{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	orderItems, err := o.orderItemRepository.FindOrderItems(ctx, id)
	if err != nil {
		return order.OrderDetail{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return order.OrderDetail{
		Order: order.Order{
			ID:         orderEntity.ID,
			UserID:     orderEntity.UserID,
			TotalPrice: orderEntity.TotalPrice,
			Status:     orderEntity.Status,
		},
		Items: orderItems,
	}, nil
}

func (o *orderServiceImpl) GetOrderByUserId(ctx context.Context, userId int) ([]order.OrderDetail, *helper.AppError) {
	orders, err := o.orderRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	var results []order.OrderDetail
	for _, item := range orders {
		orderItems, err := o.orderItemRepository.FindOrderItems(ctx, item.ID)
		if err != nil {
			return nil, helper.NewAppError(
				http.StatusInternalServerError,
				"Internal Server Error",
				err,
			)
		}

		orderDetail := order.OrderDetail{
			Order: order.Order{
				ID:         item.ID,
				UserID:     item.UserID,
				TotalPrice: item.TotalPrice,
				Status:     item.Status,
			},
			Items: orderItems,
		}

		results = append(results, orderDetail)
	}

	return results, nil
}

func (o *orderServiceImpl) UpdateOrderStatus(ctx context.Context, id int, status order.Status) *helper.AppError {
	_, err := o.orderRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	if err := o.orderRepository.UpdateStatus(ctx, id, status); err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return nil
}

func (o *orderServiceImpl) CancelOrder(ctx context.Context, id int) *helper.AppError {
	orderEntity, err := o.orderRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	if orderEntity.Status != order.StatusPending {
		return helper.NewAppError(
			http.StatusConflict,
			"Order Cannot Be Cancelled",
			errors.New("Only pending orders can be cancelled"),
		)
	}

	if err := o.orderRepository.UpdateStatus(ctx, orderEntity.ID, orderEntity.Status); err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return nil
}
