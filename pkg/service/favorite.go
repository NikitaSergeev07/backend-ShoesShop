package service

import (
	"ShoesShop"
	"ShoesShop/pkg/repository"
	"fmt"
)

type FavoriteService struct {
	repo repository.FavoriteRepository
}

func NewFavoriteService(repo repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) AddFavorite(favorite ShoesShop.Favorite) (int, error) {
	id, err := s.repo.AddFavorite(favorite)
	if err != nil {
		return 0, fmt.Errorf("failed to add favorite: %w", err)
	}
	return id, nil
}

func (s *FavoriteService) RemoveFavorite(userId, itemId int) error {
	err := s.repo.RemoveFavorite(userId, itemId)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}
	return nil
}

func (s *FavoriteService) GetFavoritesByUserId(userId int) ([]ShoesShop.Item, error) {
	items, err := s.repo.GetFavoritesByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorites for user %d: %w", userId, err)
	}
	return items, nil
}
