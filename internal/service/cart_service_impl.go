package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/cart"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type cartServiceImpl struct {
	tx                 *helper.Transaction
	cartRepository     cart.Repository
	cartItemRepository cart.ItemRepository
}

func NewCartServiceImpl(tx *helper.Transaction, cartRepository cart.Repository, cartItemRepository cart.ItemRepository) cart.Service {
	return &cartServiceImpl{tx: tx, cartRepository: cartRepository, cartItemRepository: cartItemRepository}
}

func (c *cartServiceImpl) GetItems(ctx context.Context, userId int) ([]cart.Item, *helper.AppError) {
	var cartItems []cart.Item

	err := func() error {
		cartData, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		cartItems, err = c.cartItemRepository.FindAllByCartId(ctx, cartData.ID)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, helper.ErrCartNotFound) {
			return nil, nil
		}
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return cartItems, nil
}

func (c *cartServiceImpl) AddItem(ctx context.Context, userId int, productId string, quantity int) (cart.Item, *helper.AppError) {
	var result *cart.Item

	err := c.tx.ExecTx(ctx, func(ctx context.Context) error {
		cartData, err := c.cartRepository.FindOrCreateByUserId(ctx, userId)
		if err != nil {
			return err
		}

		cartItem, err := c.cartItemRepository.FindByCartAndProductId(ctx, cartData.ID, productId)
		if err != nil {
			return err
		}

		if cartItem != nil {
			cartItem.Quantity += quantity

			updateCartItem := cart.UpdateItem{
				ID:       cartItem.ID,
				Quantity: cartItem.Quantity,
			}

			err := c.cartItemRepository.Update(ctx, updateCartItem)
			if err != nil {
				return err
			}

			result = cartItem

			return nil
		}

		result = &cart.Item{
			CartID:    cartData.ID,
			ProductID: productId,
			Quantity:  quantity,
		}
		err = c.cartItemRepository.Create(ctx, result)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, helper.ErrCartItemNotFound) {
			return cart.Item{}, helper.NewAppError(
				http.StatusNotFound,
				"Cart Item Not Found",
				err,
			)
		}

		return cart.Item{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return *result, nil
}

func (c *cartServiceImpl) UpdateItemQuantity(ctx context.Context, userId int, updateItem cart.UpdateItem) *helper.AppError {
	err := func() error {
		cartItem, err := c.cartItemRepository.FindById(ctx, updateItem.ID)
		if err != nil {
			return err
		}

		cartData, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		if cartItem.CartID != cartData.ID {
			return helper.ErrCartItemNotFound
		}

		updateItem.Quantity += cartItem.Quantity
		err = c.cartItemRepository.Update(ctx, updateItem)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, helper.ErrCartItemNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Cart Item Not Found",
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

func (c *cartServiceImpl) DeleteItem(ctx context.Context, userId int, itemId int) *helper.AppError {
	err := func() error {
		cartItem, err := c.cartItemRepository.FindById(ctx, itemId)
		if err != nil {
			return err
		}

		cartData, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		if cartItem.CartID != cartData.ID {
			return helper.ErrCartItemNotFound
		}

		err = c.cartItemRepository.Delete(ctx, itemId)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, helper.ErrCartItemNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Cart Item Not Found",
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
