package ShoesShop

type Cart struct {
	Id     int    `json:"id" db:"id"`
	UserId int    `json:"user_id" db:"user_id"` // Внешний ключ на пользователя
	User   *User  `json:"user" db:"-"`          // Поле для джойна (если нужно)
	ItemId int    `json:"item_id" db:"item_id"` // Внешний ключ на продукт
	Item   *Item  `json:"item" db:"-"`          // Поле для джойна с продуктом
	Size   string `json:"size" db:"size"`       // Поле для хранения выбранного размера
}
