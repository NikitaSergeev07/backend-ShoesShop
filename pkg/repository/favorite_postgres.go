package repository

import (
	"ShoesShop"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type FavoritePostgres struct {
	db *sqlx.DB
}

func NewFavoritePostgres(db *sqlx.DB) *FavoritePostgres {
	return &FavoritePostgres{db: db}
}

func (r *FavoritePostgres) AddFavorite(f ShoesShop.Favorite) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, item_id) VALUES ($1, $2) RETURNING id",
		FavoritesTable,
	)
	row := r.db.QueryRow(query, f.UserId, f.ItemId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to add favorite: %w", err)
	}

	return id, nil
}

func (r *FavoritePostgres) RemoveFavorite(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND item_id = $2", FavoritesTable)

	_, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}

	return nil
}

func (r *FavoritePostgres) GetFavoritesByUserId(userId int) ([]ShoesShop.Item, error) {
	var items []ShoesShop.Item
	query := fmt.Sprintf(`
		SELECT i.id, i.title, i.price, i.description, i.image_urls
		FROM %s f
		INNER JOIN %s i ON f.item_id = i.id
		WHERE f.user_id = $1`,
		FavoritesTable, ItemsTable,
	)

	rows, err := r.db.Queryx(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorites: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item ShoesShop.Item
		var imageURLs pq.StringArray

		if err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Description, &imageURLs); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}

		item.ImageURLs = imageURLs
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return items, nil
}
