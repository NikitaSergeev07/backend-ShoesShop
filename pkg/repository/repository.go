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

type FavoriteRepository interface {
	AddFavorite(f ShoesShop.Favorite) (int, error)
	RemoveFavorite(userId, itemId int) error
	GetFavoritesByUserId(userId int) ([]ShoesShop.Item, error)
}

type CartRepository interface {
	AddToCart(cart ShoesShop.Cart) (int, error)
	RemoveFromCart(userId, itemId int, size string) error
	GetCartByUserId(userId int) ([]ShoesShop.Cart, error)
	RemoveAll(userId int) error
}

type Repository struct {
	Authorization
	ItemRepository
	ReviewRepository
	FavoriteRepository
	CartRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:      NewAuthPostgres(db),
		ItemRepository:     NewItemPostgres(db),
		ReviewRepository:   NewReviewPostgres(db),
		FavoriteRepository: NewFavoritePostgres(db),
		CartRepository:     NewCartPostgres(db),
	}
}
