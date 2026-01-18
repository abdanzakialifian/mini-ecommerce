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
	cartRepository     cart.CartRepository
	cartItemRepository cart.CartItemRepository
}

func NewCartServiceImpl(tx *helper.Transaction, cartRepository cart.CartRepository, cartItemRepository cart.CartItemRepository) cart.CartService {
	return &cartServiceImpl{tx: tx, cartRepository: cartRepository, cartItemRepository: cartItemRepository}
}

func (c *cartServiceImpl) GetCartItems(ctx context.Context, userId int) ([]cart.CartItem, *helper.AppError) {
	var results []cart.CartItem

	err := func() error {
		cart, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		results, err = c.cartItemRepository.FindAllByCartId(ctx, cart.ID)
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

	return results, nil
}

func (c *cartServiceImpl) AddCartItemToCart(ctx context.Context, userId int, productId string, quantity int) (cart.CartItem, *helper.AppError) {
	var result *cart.CartItem

	err := c.tx.ExecTx(ctx, func(ctx context.Context) error {
		userCart, err := c.cartRepository.FindOrCreateByUserId(ctx, userId)
		if err != nil {
			return err
		}

		cartItem, err := c.cartItemRepository.FindByCartAndProductId(ctx, userCart.ID, productId)
		if err != nil {
			return err
		}

		if cartItem != nil {
			cartItem.Quantity += quantity

			updateCartItem := cart.UpdateCartItem{
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

		result = &cart.CartItem{
			CartID:    userCart.ID,
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
			return cart.CartItem{}, helper.NewAppError(
				http.StatusNotFound,
				"Cart Item Not Found",
				err,
			)
		}

		return cart.CartItem{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return *result, nil
}

func (c *cartServiceImpl) UpdateCartItemQuantity(ctx context.Context, userId int, updateCartItem cart.UpdateCartItem) *helper.AppError {
	err := func() error {
		cartItem, err := c.cartItemRepository.FindById(ctx, updateCartItem.ID)
		if err != nil {
			return err
		}

		cart, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		if cartItem.CartID != cart.ID {
			return helper.ErrCartItemNotFound
		}

		updateCartItem.Quantity += cartItem.Quantity
		err = c.cartItemRepository.Update(ctx, updateCartItem)
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

func (c *cartServiceImpl) DeleteCartItemFromCart(ctx context.Context, userId int, cartItemId int) *helper.AppError {
	err := func() error {
		cartItem, err := c.cartItemRepository.FindById(ctx, cartItemId)
		if err != nil {
			return err
		}

		cart, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		if cartItem.CartID != cart.ID {
			return helper.ErrCartItemNotFound
		}

		err = c.cartItemRepository.Delete(ctx, cartItemId)
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
