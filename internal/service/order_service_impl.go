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
	orderRepository     order.Repository
	orderItemRepository order.ItemRepository
	productRepository   product.Repository
}

func NewOrder(tx *helper.Transaction, orderRepository order.Repository, orderItemRepository order.ItemRepository) order.Service {
	return &orderServiceImpl{tx: tx, orderRepository: orderRepository, orderItemRepository: orderItemRepository}
}

func (o *orderServiceImpl) Create(ctx context.Context, userId int, newItems []order.NewItem) (order.Detail, *helper.AppError) {
	var orderDetail order.Detail
	err := o.tx.ExecTx(ctx, func(ctx context.Context) error {
		var totalPrice float64

		var orderItems []order.Item
		for _, newItem := range newItems {
			productData, err := o.productRepository.Find(ctx, newItem.ProductID)
			if err != nil {
				return err
			}

			if productData.Stock < newItem.Quantity {
				return helper.ErrProductInsufficientStock
			}

			itemPrice := productData.Price * float64(newItem.Quantity)
			totalPrice += itemPrice

			orderItem := order.Item{
				ProductID: newItem.ProductID,
				Price:     productData.Price,
				Quantity:  newItem.Quantity,
			}

			orderItems = append(orderItems, orderItem)
		}

		orderData := order.Data{
			UserID:     userId,
			TotalPrice: totalPrice,
			Status:     order.StatusPending,
		}

		if err := o.orderRepository.Create(ctx, &orderData); err != nil {
			return err
		}

		for i := range orderItems {
			orderItems[i].OrderID = orderData.ID
		}

		if err := o.orderItemRepository.CreateItems(ctx, orderItems); err != nil {
			return err
		}

		for _, orderItem := range orderItems {
			if err := o.productRepository.UpdateStock(ctx, orderItem.ProductID, orderItem.Quantity); err != nil {
				return err
			}
		}

		orderDetail = order.Detail{
			Data: order.Data{
				ID:         orderData.ID,
				UserID:     orderData.UserID,
				TotalPrice: orderData.TotalPrice,
				Status:     orderData.Status,
			},
			Items: orderItems,
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

func (o *orderServiceImpl) Get(ctx context.Context, id int) (order.Detail, *helper.AppError) {
	orderData, err := o.orderRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrOrderNotFound) {
			return order.Detail{}, helper.NewAppError(
				http.StatusNotFound,
				"Order Not Found",
				err,
			)
		}

		return order.Detail{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	orderItems, err := o.orderItemRepository.FindItems(ctx, id)
	if err != nil {
		return order.Detail{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return order.Detail{
		Data: order.Data{
			ID:         orderData.ID,
			UserID:     orderData.UserID,
			TotalPrice: orderData.TotalPrice,
			Status:     orderData.Status,
		},
		Items: orderItems,
	}, nil
}

func (o *orderServiceImpl) GetByUserId(ctx context.Context, userId int) ([]order.Detail, *helper.AppError) {
	orders, err := o.orderRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	var orderDetails []order.Detail
	for _, orderData := range orders {
		orderItems, err := o.orderItemRepository.FindItems(ctx, orderData.ID)
		if err != nil {
			return nil, helper.NewAppError(
				http.StatusInternalServerError,
				"Internal Server Error",
				err,
			)
		}

		orderDetail := order.Detail{
			Data: order.Data{
				ID:         orderData.ID,
				UserID:     orderData.UserID,
				TotalPrice: orderData.TotalPrice,
				Status:     orderData.Status,
			},
			Items: orderItems,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}

func (o *orderServiceImpl) UpdateStatus(ctx context.Context, id int, status order.Status) *helper.AppError {
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

func (o *orderServiceImpl) Cancel(ctx context.Context, id int) *helper.AppError {
	orderData, err := o.orderRepository.FindById(ctx, id)
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

	if orderData.Status != order.StatusPending {
		return helper.NewAppError(
			http.StatusConflict,
			"Order Cannot Be Cancelled",
			errors.New("Only pending orders can be cancelled"),
		)
	}

	if err := o.orderRepository.UpdateStatus(ctx, orderData.ID, orderData.Status); err != nil {
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
