package service

import (
	"fmt"

	"frappuccino/internal/dal"
	"frappuccino/models"
)

func isUniqueMenu(items []models.MenuItem, id string) bool {
	for _, item := range items {
		if item.ID == id {
			return false
		}
	}
	return true
}

func AddMenuItem(items []models.MenuItem, newItem models.MenuItem) Status {
	if newItem.ID == "" {
		return NoIdGiven
	}
	if !isUniqueMenu(items, newItem.ID) {
		return IdAlreadyExists
	}
	items = append(items, newItem)
	err := dal.WriteMenu(items)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func UpdateMenuItem(items []models.MenuItem, newItem models.MenuItem, id string) ([]models.MenuItem, Status) {
	if isUniqueMenu(items, id) {
		return nil, NotFound
	}
	if id != newItem.ID {
		return nil, Status{fmt.Errorf("%s should equal to id in body.", id), 400}
	}
	newItems := make([]models.MenuItem, 0)

	for _, item := range items {
		if item.ID == id {
			newItems = append(newItems, newItem)
			continue
		}
		newItems = append(newItems, item)
	}
	err := dal.WriteMenu(newItems)
	if err != nil {
		return nil, Status{err, 500}
	}
	return newItems, Success
}

func FindMenuItem(items []models.MenuItem, id string) (models.MenuItem, Status) {
	for _, item := range items {
		if item.ID == id {
			return item, Success
		}
	}
	return models.MenuItem{}, NotFound
}

func RemoveMenuItem(items []models.MenuItem, id string) Status {
	if isUniqueMenu(items, id) {
		return NotFound
	}
	newItems := make([]models.MenuItem, 0)

	for _, item := range items {
		if item.ID == id {
			continue
		}
		newItems = append(newItems, item)
	}
	err := dal.WriteMenu(newItems)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}
