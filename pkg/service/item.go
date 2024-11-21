package service

import (
	"ShoesShop"
	"ShoesShop/enums"
	"ShoesShop/pkg/repository"
	"fmt"
)

type ItemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item ShoesShop.Item) (int, error) {
	id, err := s.repo.CreateItem(item)
	if err != nil {
		return 0, fmt.Errorf("failed to create item: %w", err)
	}
	return id, nil
}

func (s *ItemService) SearchItems(query string, option enums.SortOption) ([]ShoesShop.Item, error) {
	items, err := s.repo.SearchItems(query, option)
	if err != nil {
		return nil, fmt.Errorf("failed to search items: %w", err)
	}

	return items, nil
}

func (s *ItemService) GetItemById(id int) (ShoesShop.Item, error) {
	item, err := s.repo.GetItemById(id)
	if err != nil {
		return item, fmt.Errorf("failed to get item: %w", err)
	}
	return item, nil
}

func (s *ItemService) GetAllItems() ([]ShoesShop.Item, error) {
	items, err := s.repo.GetAllItems()
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}
	return items, nil
}

func (s *ItemService) UpdateItem(item ShoesShop.Item) error {
	err := s.repo.UpdateItem(item)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (s *ItemService) DeleteItem(id int) error {
	err := s.repo.DeleteItem(id)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}
	return nil
}
