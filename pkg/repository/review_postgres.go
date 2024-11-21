package repository

import (
	"ShoesShop"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (r *ReviewPostgres) CreateReview(review ShoesShop.Review) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name_of_reviewer, text, score, category, date, user_id, item_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		ReviewsTable,
	)
	row := r.db.QueryRow(query, review.NameOfReviewer, review.Text, review.Score, review.Category, review.Date, review.UserId, review.ItemId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create review: %w", err)
	}

	return id, nil
}

func (r *ReviewPostgres) GetAllReviews() ([]ShoesShop.Review, error) {
	query := fmt.Sprintf("SELECT id, name_of_reviewer, text, score, category, date, user_id, item_id FROM %s", ReviewsTable)

	var reviews []ShoesShop.Review
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all reviews: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var review ShoesShop.Review
		if err := rows.StructScan(&review); err != nil {
			return nil, fmt.Errorf("failed to scan review: %w", err)
		}
		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return reviews, nil
}
