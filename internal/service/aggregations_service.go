package service

import (
	"frappuccino/internal/dal"
	"frappuccino/models"
)

func SortByPopularity(orders []models.Order) ([]models.MenuItem, Status) {
	items, err := dal.ReadMenu()
	if err != nil {
		return nil, Status{err, 500}
	}
	var sortedItems []models.MenuItem
	keys := dal.SortItemID(orders)
	for _, key := range keys {
		temp, status := FindMenuItem(items, key)
		if status.ErrorMessage != nil {
			return nil, status
		}
		sortedItems = append(sortedItems, temp)
	}
	return sortedItems, Status{nil, 200}
}

func SumSales(orders []models.Order) (float64, Status) {
	var sum float64
	items, err := dal.ReadMenu()
	if err != nil {
		return 0, Status{err, 500}
	}
	sum = 0
	for _, order := range orders {
		if order.Status == "closed" {
			for _, item := range order.Items {
				temp, status := FindMenuItem(items, item.ProductID)
				if status.ErrorMessage != nil {
					return 0, status
				}
				sum += (float64(item.Quantity) * temp.Price)
			}
		}
	}
	return sum, Status{nil, 200}
}
