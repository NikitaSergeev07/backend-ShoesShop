package ShoesShop

type User struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" binding:"required" db:"email"`
	PhoneNumber string `json:"phone_number" binding:"required" db:"phone_number"`
	Password    string `json:"password" binding:"required" db:"password_hash"`
}
