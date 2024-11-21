package repository

import (
	"ShoesShop"
	"ShoesShop/enums"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ShoesShop.User) (int, error)
	GetUser(email, password string) (ShoesShop.User, error)
}

type ItemRepository interface {
	CreateItem(item ShoesShop.Item) (int, error)
	GetAllItems() ([]ShoesShop.Item, error)
	SearchItems(query string, sortOption enums.SortOption) ([]ShoesShop.Item, error)
	GetItemById(id int) (ShoesShop.Item, error)
	UpdateItem(item ShoesShop.Item) error
	DeleteItem(id int) error
}

type ReviewRepository interface {
	CreateReview(review ShoesShop.Review) (int, error)
	GetAllReviews() ([]ShoesShop.Review, error)
}

type Repository struct {
	Authorization
	ItemRepository
	ReviewRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:    NewAuthPostgres(db),
		ItemRepository:   NewItemPostgres(db),
		ReviewRepository: NewReviewPostgres(db),
	}
}
