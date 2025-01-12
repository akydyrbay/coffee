package service

import (
	"fmt"
	"time"

	"frappuccino/internal/dal"
	"frappuccino/models"
)

// NOTE: Service is about how data operates, not how it looks. Should include every operations with orders subfunctions: order.({data})
func isUniqueOrder(orders []models.Order, id string) bool {
	for _, order := range orders {
		if order.ID == id {
			return false
		}
	}
	return true
}

func completeOrder(newOrder models.Order) Status {
	// CHECK.
	inventory, err := dal.ReadInventory()
	if err != nil {
		return Status{ErrorMessage: err, Code: 500}
	}
	menu, err := dal.ReadMenu()
	if err != nil {
		return Status{ErrorMessage: err, Code: 500}
	}
	err = checkOrder(newOrder, menu, inventory)
	if err != nil {
		return Status{ErrorMessage: err, Code: 400}
	}
	// DEDUCT.
	err = updateQuantity(newOrder, menu, inventory)
	if err != nil {
		return Status{ErrorMessage: err, Code: 500}
	}
	return Success
}

func AddOrder(orders []models.Order, newOrder models.Order) Status {
	if newOrder.ID == "" {
		return NoIdGiven
	}
	if !isUniqueOrder(orders, newOrder.ID) {
		return IdAlreadyExists
	}
	newOrder.Status = "open"
	t := time.Now()
	t = t.UTC()
	newOrder.CreatedAt = t.Format("2006-01-02T15:04:05Z07:00")
	orders = append(orders, newOrder)
	err := dal.WriteOrders(orders)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func FindOrder(orders []models.Order, id string) (models.Order, Status) {
	for _, order := range orders {
		if order.ID == id {
			return order, Success
		}
	}
	return models.Order{}, NotFound
}

func UpdateOrder(orders []models.Order, newOrder models.Order, id string) Status {
	if isUniqueOrder(orders, id) {
		return NotFound
	}
	if id != newOrder.ID {
		return Status{fmt.Errorf("%s should equal to id in body.", id), 400}
	}

	order, response := FindOrder(orders, id)
	if response.ErrorMessage != nil {
		return response
	}
	if order.Status != "open" {
		return Status{fmt.Errorf("Order is not opened."), 400}
	}
	newOrder.Status = "open"
	newOrders := make([]models.Order, 0)
	for _, order := range orders {
		if order.ID == id {
			newOrders = append(newOrders, newOrder)
			continue
		}
		newOrders = append(newOrders, order)
	}
	err := dal.WriteOrders(newOrders)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func RemoveOrder(orders []models.Order, id string) Status {
	if isUniqueOrder(orders, id) {
		return NotFound
	}
	newOrders := make([]models.Order, 0)
	for _, order := range orders {
		if order.ID == id {
			continue
		}
		newOrders = append(newOrders, order)
	}
	err := dal.WriteOrders(newOrders)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func FinishOrder(orders []models.Order, id string, status string) Status {
	// if isUniqueOrder(orders, id) {
	// 	return NotFound
	// }
	newOrder, response := FindOrder(orders, id)
	if response.ErrorMessage != nil {
		return response
	}
	if newOrder.Status != "open" {
		return Status{fmt.Errorf("The order is already closed."), 400}
	}
	response = completeOrder(newOrder)
	if response.ErrorMessage != nil {
		return response
	}
	newOrder.Status = status

	newOrders := make([]models.Order, 0)
	for _, order := range orders {
		if order.ID == id {
			newOrders = append(newOrders, newOrder)
			continue
		}
		newOrders = append(newOrders, order)
	}

	err := dal.WriteOrders(newOrders)
	if err != nil {
		return Status{err, 500}
	}
	return Success
}

func checkOrder(order models.Order, menuItems []models.MenuItem, inventoryItems []models.InventoryItem) error {
	for _, orderItem := range order.Items {
		menuItem, response := FindMenuItem(menuItems, orderItem.ProductID)
		if response == NotFound {
			return fmt.Errorf("Invalid product ID in order items.")
		}

		if response.ErrorMessage != nil {
			return response.ErrorMessage
		}
		for _, orderIngridient := range menuItem.Ingredients {
			inventoryItem, response := FindInventoryItem(inventoryItems, orderIngridient.IngredientID)
			if response.ErrorMessage != nil {
				return response.ErrorMessage
			}
			if inventoryItem.Quantity-(orderIngridient.Quantity*orderItem.Quantity) < 0 {
				return fmt.Errorf("Insufficient inventory for ingredient '%s'. Required: %d%s, Available: %d%s.", orderIngridient.IngredientID, orderIngridient.Quantity*orderItem.Quantity, inventoryItem.Unit, inventoryItem.Quantity, inventoryItem.Unit)
			}
		}
	}
	return nil
}

func updateQuantity(order models.Order, menuItems []models.MenuItem, inventoryItems []models.InventoryItem) error {
	newInventory, err := dal.ReadInventory()
	if err != nil {
		return err
	}
	for _, orderItem := range order.Items {
		menuItem, response := FindMenuItem(menuItems, orderItem.ProductID)
		if response.ErrorMessage != nil {
			return response.ErrorMessage
		}
		for _, orderIngridient := range menuItem.Ingredients {
			newInventoryItem, response := FindInventoryItem(inventoryItems, orderIngridient.IngredientID)
			if response.ErrorMessage != nil {
				return response.ErrorMessage
			}
			newInventoryItem.Quantity -= (orderIngridient.Quantity * orderItem.Quantity)
			newInventory, response = UpdateInventoryItem(newInventory, newInventoryItem, newInventoryItem.IngredientID)
			err = dal.WriteInventory(newInventory)
			if err != nil {
				return err
			}
			if response.ErrorMessage != nil {
				return response.ErrorMessage
			}
		}
	}
	return nil
}
