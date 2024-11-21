package service

import (
	"ShoesShop"
	"ShoesShop/enums"
	"ShoesShop/pkg/repository"
	"fmt"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(item ShoesShop.Review) (int, error) {
	if item.Category == enums.Product && (item.ItemId == nil || *item.ItemId == 0) {
		return 0, fmt.Errorf("item_id is required for category 'Product'")
	}

	id, err := s.repo.CreateReview(item)
	if err != nil {
		return 0, fmt.Errorf("failed to create review: %w", err)
	}
	return id, nil
}

func (s *ReviewService) GetAllReviews() ([]ShoesShop.Review, error) {
	reviews, err := s.repo.GetAllReviews()
	if err != nil {
		return nil, fmt.Errorf("failed to get all reviews: %w", err)
	}
	return reviews, nil
}
