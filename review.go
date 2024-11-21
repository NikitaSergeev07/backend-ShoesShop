package ShoesShop

import (
	"ShoesShop/enums"
	"time"
)

type Review struct {
	Id             int                  `json:"id" db:"id"`
	NameOfReviewer string               `json:"name_of_reviewer" binding:"required" db:"name_of_reviewer"`
	Text           string               `json:"text" binding:"required" db:"text"`
	Score          int                  `json:"score" binding:"required" db:"score"`
	Category       enums.CategoryOption `json:"category" binding:"required" db:"category"`
	Date           time.Time            `json:"date" binding:"required" db:"date"`

	UserId int   `json:"user_id" db:"user_id"` // Внешний ключ на пользователя
	User   *User `json:"user" db:"-"`          // Поле для джойна (если нужно)

	ItemId *int  `json:"item_id" db:"item_id"` // Внешний ключ на продукт (необязательное поле)
	Item   *Item `json:"item" db:"-"`          // Поле для джойна с продуктом
}
