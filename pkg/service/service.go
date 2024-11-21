package service

import (
	"ShoesShop"
	"ShoesShop/enums"
	"ShoesShop/pkg/repository"
)

type Authorization interface {
	CreateUser(user ShoesShop.User) (int, error)
	GetUser(email, password string) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	InvalidateToken(token string) error
	IsTokenBlacklisted(token string) bool
}

type Item interface {
	CreateItem(item ShoesShop.Item) (int, error)
	GetItemById(id int) (ShoesShop.Item, error)
	SearchItems(query string, option enums.SortOption) ([]ShoesShop.Item, error)
	GetAllItems() ([]ShoesShop.Item, error)
	UpdateItem(item ShoesShop.Item) error
	DeleteItem(id int) error
}

type Review interface {
	CreateReview(review ShoesShop.Review) (int, error)
	GetAllReviews() ([]ShoesShop.Review, error)
}

type Service struct {
	Authorization
	Item
	Review
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Item:          NewItemService(repos.ItemRepository),
		Review:        NewReviewService(repos.ReviewRepository),
	}
}
