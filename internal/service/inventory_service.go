package service

import (
	"fmt"

	"frappuccino/internal/dal"
	"frappuccino/models"
)

// NOTE: Service is about how data operates () not how it looks. Should include every operations with item subfunctions: item.({data})
// CreateInventoryItem Creates ...

func isUniqueInventory(items []models.InventoryItem, id string) bool {
	for _, item := range items {
		if item.IngredientID == id {
			return false
		}
	}
	return true
}

func AddInventoryItem(items []models.InventoryItem, newItem models.InventoryItem) Status {
	if newItem.IngredientID == "" {
		return NoIdGiven
	}
	if !isUniqueInventory(items, newItem.IngredientID) {
		return IdAlreadyExists
	}
	items = append(items, newItem)
	err := dal.WriteInventory(items)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func FindInventoryItem(items []models.InventoryItem, id string) (models.InventoryItem, Status) {
	for i, item := range items {
		if item.IngredientID == id {
			return items[i], Success
		}
	}
	return models.InventoryItem{}, NotFound
}

func UpdateInventoryItem(items []models.InventoryItem, newItem models.InventoryItem, id string) ([]models.InventoryItem, Status) {
	if isUniqueInventory(items, id) {
		return nil, NotFound
	}
	if id != newItem.IngredientID {
		return nil, Status{fmt.Errorf("%s should equal to id in body.", id), 400}
	}
	newItems := make([]models.InventoryItem, 0)

	for _, item := range items {
		if item.IngredientID == id {
			newItems = append(newItems, newItem)
			continue
		}
		newItems = append(newItems, item)
	}
	err := dal.WriteInventory(newItems)
	if err != nil {
		return nil, Status{err, 500}
	}
	return newItems, Success
}

func RemoveInventoryItem(items []models.InventoryItem, id string) Status {
	if isUniqueInventory(items, id) {
		return NotFound
	}
	newItems := make([]models.InventoryItem, 0)

	for _, item := range items {
		if item.IngredientID == id {
			continue
		}
		newItems = append(newItems, item)
	}
	err := dal.WriteInventory(newItems)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}
