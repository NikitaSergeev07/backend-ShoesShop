package service

import (
	"ShoesShop"
	"ShoesShop/enums"
	"ShoesShop/pkg/repository"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
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

type Favorite interface {
	AddFavorite(f ShoesShop.Favorite) (int, error)
	RemoveFavorite(userId, itemId int) error
	GetFavoritesByUserId(userId int) ([]ShoesShop.Item, error)
}

type Cart interface {
	AddToCart(cart ShoesShop.Cart) (int, error)
	RemoveFromCart(userId, itemId int, size string) error
	GetCartByUserId(userId int) ([]ShoesShop.Cart, error)
	RemoveAll(userId int) error
}

type Payment interface {
	CreatePayment(amount, description string) (*yoopayment.Payment, error)
}

type Service struct {
	Authorization
	Item
	Review
	Favorite
	Cart
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Item:          NewItemService(repos.ItemRepository),
		Review:        NewReviewService(repos.ReviewRepository),
		Favorite:      NewFavoriteService(repos.FavoriteRepository),
		Cart:          NewCartService(repos.CartRepository),
		Payment:       NewPaymentService("995752", "test_BonXu0DzzvUyHD7gLcQXETBKW-AOggGeyNvSQwaYZ8U"),
	}
}
