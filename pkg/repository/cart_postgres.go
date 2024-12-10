package repository

import (
	"ShoesShop"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (r *CartPostgres) AddToCart(cart ShoesShop.Cart) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, item_id, size) VALUES ($1, $2, $3) RETURNING id",
		CartsTable,
	)
	row := r.db.QueryRow(query, cart.UserId, cart.ItemId, cart.Size)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to add item to cart: %w", err)
	}

	return id, nil
}

func (r *CartPostgres) RemoveFromCart(userId, itemId int, size string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND item_id = $2 AND size = $3", CartsTable)

	_, err := r.db.Exec(query, userId, itemId, size)
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %w", err)
	}

	return nil
}

func (r *CartPostgres) GetCartByUserId(userId int) ([]ShoesShop.Cart, error) {
	var cartItems []ShoesShop.Cart
	query := fmt.Sprintf(`
		SELECT ci.id, ci.user_id, ci.item_id, ci.size, 
		       i.title, i.price, i.description, i.image_urls
		FROM %s ci
		INNER JOIN %s i ON ci.item_id = i.id
		WHERE ci.user_id = $1`,
		CartsTable, ItemsTable,
	)

	rows, err := r.db.Queryx(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem ShoesShop.Cart
		var item ShoesShop.Item
		var imageURLs pq.StringArray

		if err := rows.Scan(&cartItem.Id, &cartItem.UserId, &cartItem.ItemId, &cartItem.Size,
			&item.Title, &item.Price, &item.Description, &imageURLs); err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}

		item.ImageURLs = imageURLs
		cartItem.Item = &item

		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return cartItems, nil
}

func (r *CartPostgres) RemoveAll(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", CartsTable)

	_, err := r.db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("failed to remove all items from cart for user %d: %w", userId, err)
	}

	return nil
}
