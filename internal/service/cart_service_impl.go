package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type cartServiceImpl struct {
	tx                 *helper.Transaction
	cartRepository     domain.CartRepository
	cartItemRepository domain.CartItemRepository
}

func NewCartServiceImpl(tx *helper.Transaction, cartRepository domain.CartRepository, cartItemRepository domain.CartItemRepository) domain.CartService {
	return &cartServiceImpl{tx: tx, cartRepository: cartRepository, cartItemRepository: cartItemRepository}
}

func (c *cartServiceImpl) GetCartItems(ctx context.Context, userId int) ([]model.CartItem, *helper.AppError) {
	var cartItems []model.CartItem

	err := func() error {
		cart, err := c.cartRepository.FindByUserId(ctx, userId)
		if err != nil {
			return err
		}

		cartItems, err = c.cartItemRepository.FindAllByCartId(ctx, cart.ID)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, domain.ErrCartNotFound) {
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

func (c *cartServiceImpl) AddCartItemToCart(ctx context.Context, userId int, productId string, quantity int) (model.CartItem, *helper.AppError) {
	var cartItemModel *model.CartItem

	err := c.tx.ExecTx(ctx, func(ctx context.Context) error {
		cart, err := c.cartRepository.FindOrCreateByUserId(ctx, userId)
		if err != nil {
			return err
		}

		cartItem, err := c.cartItemRepository.FindByCartAndProductId(ctx, cart.ID, productId)
		if err != nil {
			return err
		}

		if cartItem != nil {
			cartItem.Quantity += quantity

			updateCartItem := model.UpdateCartItem{
				ID:       cartItem.ID,
				Quantity: cartItem.Quantity,
			}

			err := c.cartItemRepository.Update(ctx, updateCartItem)
			if err != nil {
				return err
			}

			cartItemModel = cartItem

			return nil
		}

		cartItemModel = &model.CartItem{
			CartID:    cart.ID,
			ProductID: productId,
			Quantity:  quantity,
		}
		err = c.cartItemRepository.Create(ctx, cartItemModel)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
			return model.CartItem{}, helper.NewAppError(
				http.StatusNotFound,
				"Cart Item Not Found",
				err,
			)
		}

		return model.CartItem{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return *cartItemModel, nil
}

func (c *cartServiceImpl) UpdateCartItemQuantity(ctx context.Context, userId int, updateCartItem model.UpdateCartItem) *helper.AppError {
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
			return domain.ErrCartItemNotFound
		}

		updateCartItem.Quantity += cartItem.Quantity
		err = c.cartItemRepository.Update(ctx, updateCartItem)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
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
			return domain.ErrCartItemNotFound
		}

		err = c.cartItemRepository.Delete(ctx, cartItemId)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
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
