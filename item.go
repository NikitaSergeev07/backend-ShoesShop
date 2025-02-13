package ShoesShop

type Item struct {
	Id          int      `json:"id" db:"id"`
	Title       string   `json:"title" binding:"required" db:"title"`
	Price       float64  `json:"price" binding:"required" db:"price"`
	Description string   `json:"description" binding:"required" db:"description"`
	ImageURLs   []string `json:"image_urls" db:"image_urls"`
	Category    string   `json:"category" db:"category"`
}
