package service

import (
	"ShoesShop"
	"ShoesShop/pkg/repository"
	"fmt"
)

type CartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(cart ShoesShop.Cart) (int, error) {
	existingCart, err := s.repo.GetCartByUserId(cart.UserId)
	if err != nil {
		return 0, fmt.Errorf("failed to get cart items: %w", err)
	}
	for _, item := range existingCart {
		if item.ItemId == cart.ItemId && item.Size == cart.Size {
			return 0, fmt.Errorf("item with size %s already in cart", cart.Size)
		}
	}

	id, err := s.repo.AddToCart(cart)
	if err != nil {
		return 0, fmt.Errorf("failed to add item to cart: %w", err)
	}

	return id, nil
}

func (s *CartService) RemoveFromCart(userId, itemId int, size string) error {
	err := s.repo.RemoveFromCart(userId, itemId, size)
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %w", err)
	}

	return nil
}

func (s *CartService) GetCartByUserId(userId int) ([]ShoesShop.Cart, error) {
	cartItems, err := s.repo.GetCartByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}

	return cartItems, nil
}

func (s *CartService) RemoveAll(userId int) error {
	err := s.repo.RemoveAll(userId)
	if err != nil {
		return fmt.Errorf("failed to remove all cart items: %w", err)
	}
	return nil
}
