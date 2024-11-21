package repository

import (
	"ShoesShop"
	"ShoesShop/enums"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(item ShoesShop.Item) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (title, price, description, image_urls) VALUES ($1, $2, $3, $4) RETURNING id",
		ItemsTable,
	)
	row := r.db.QueryRow(query, item.Title, item.Price, item.Description, pq.Array(item.ImageURLs))
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create item: %w", err)
	}

	return id, nil
}

func (r *ItemPostgres) SearchItems(query string, sortOption enums.SortOption) ([]ShoesShop.Item, error) {
	var items []ShoesShop.Item
	var orderClause string
	var whereClause string
	var args []interface{}

	switch sortOption {
	case enums.SortByName:
		orderClause = "ORDER BY title"
	case enums.SortByPriceAsc:
		orderClause = "ORDER BY price ASC"
	case enums.SortByPriceDesc:
		orderClause = "ORDER BY price DESC"
	default:
		return nil, fmt.Errorf("invalid sort option in repository")
	}

	if query != "" {
		whereClause = "WHERE title ILIKE $1"
		args = append(args, "%"+query+"%")
	}

	queryString := fmt.Sprintf(
		"SELECT id, title, price, description, image_urls FROM %s %s %s",
		ItemsTable,
		whereClause,
		orderClause,
	)

	rows, err := r.db.Queryx(queryString, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
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

func (r *ItemPostgres) GetItemById(id int) (ShoesShop.Item, error) {
	query := fmt.Sprintf("SELECT id, title, price, description, image_urls FROM %s WHERE id = $1", ItemsTable)

	var item ShoesShop.Item
	row := r.db.QueryRowx(query, id)

	if err := row.Scan(&item.Id, &item.Title, &item.Price, &item.Description, (*pq.StringArray)(&item.ImageURLs)); err != nil {
		return item, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (r *ItemPostgres) GetAllItems() ([]ShoesShop.Item, error) {
	query := fmt.Sprintf("SELECT id, title, price, description, image_urls FROM %s", ItemsTable)

	var items []ShoesShop.Item
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item ShoesShop.Item

		if err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Description, (*pq.StringArray)(&item.ImageURLs)); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

// TODO не проверял
func (r *ItemPostgres) UpdateItem(item ShoesShop.Item) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, price = $2, description = $3, image_urls = $4 WHERE id = $5", ItemsTable)

	_, err := r.db.Exec(query, item.Title, item.Price, item.Description, pq.Array(item.ImageURLs), item.Id)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

// TODO не проверял
func (r *ItemPostgres) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ItemsTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
