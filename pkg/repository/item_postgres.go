package repository

import (
	"ShoesShop"
	"ShoesShop/enums"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
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
		"INSERT INTO %s (title, price, description, image_urls, category) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		ItemsTable,
	)
	row := r.db.QueryRow(query, item.Title, item.Price, item.Description, pq.Array(item.ImageURLs), item.Category)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create item: %w", err)
	}

	return id, nil
}

func (r *ItemPostgres) SearchItems(query string, sortOption enums.SortOption) ([]ShoesShop.Item, error) {
	var items []ShoesShop.Item
	var orderClause string
	var whereClauses []string
	var args []interface{}

	// Определение порядка сортировки
	switch sortOption {
	case enums.SortByName:
		orderClause = "ORDER BY title"
	case enums.SortByPriceAsc:
		orderClause = "ORDER BY price ASC"
	case enums.SortByPriceDesc:
		orderClause = "ORDER BY price DESC"
	case enums.Shoes1:
		whereClauses = append(whereClauses, "category = $1")
		args = append(args, "Кроссовки")
	case enums.Shoes2:
		whereClauses = append(whereClauses, "category = $1")
		args = append(args, "Туфли")
	default:
		return nil, fmt.Errorf("invalid sort option in repository")
	}

	// Фильтрация по запросу (query) в title и category
	if query != "" {
		queryFilter := "(title ILIKE $%d OR category ILIKE $%d)"
		whereClauses = append(whereClauses, fmt.Sprintf(queryFilter, len(args)+1, len(args)+1))
		args = append(args, "%"+query+"%")
	}

	// Сборка WHERE условий
	whereClause := ""
	if len(whereClauses) > 0 {
		whereClause = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Формирование итогового запроса
	queryString := fmt.Sprintf(
		"SELECT id, title, price, description, image_urls, category FROM %s %s %s",
		ItemsTable,
		whereClause,
		orderClause,
	)

	rows, err := r.db.Queryx(queryString, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}
	defer rows.Close()

	// Обработка строк результата
	for rows.Next() {
		var item ShoesShop.Item
		var imageURLs pq.StringArray

		if err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Description, &imageURLs, &item.Category); err != nil {
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
	query := fmt.Sprintf("SELECT id, title, price, description, image_urls, category FROM %s WHERE id = $1", ItemsTable)

	var item ShoesShop.Item
	row := r.db.QueryRowx(query, id)

	if err := row.Scan(&item.Id, &item.Title, &item.Price, &item.Description, (*pq.StringArray)(&item.ImageURLs), &item.Category); err != nil {
		return item, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (r *ItemPostgres) GetAllItems() ([]ShoesShop.Item, error) {
	query := fmt.Sprintf("SELECT id, title, price, description, image_urls, category FROM %s", ItemsTable)

	var items []ShoesShop.Item
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item ShoesShop.Item

		if err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Description, (*pq.StringArray)(&item.ImageURLs), &item.Category); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ItemPostgres) UpdateItem(item ShoesShop.Item) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, price = $2, description = $3, image_urls = $4, category = $5 WHERE id = $6", ItemsTable)

	_, err := r.db.Exec(query, item.Title, item.Price, item.Description, pq.Array(item.ImageURLs), item.Category, item.Id)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

func (r *ItemPostgres) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ItemsTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
