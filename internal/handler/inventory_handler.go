package handler

import (
	"net/http"

	"frappuccino/config"
	"frappuccino/internal/dal"
	"frappuccino/internal/service"
)

// NOTE: Handler is about how data looks not how it operates
// TODO: fix naming

func PostInventoryItem(w http.ResponseWriter, r *http.Request) {
	items, err := dal.ReadInventory()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	newItem, err := dal.ContentToInventoryItem(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.AddInventoryItem(items, newItem)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, newItem, "Post Inventory.", 201)
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
	items, err := dal.ReadInventory()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	config.SendResponse(w, r, items, "Get Inventory.", 200)
}

func GetInventoryItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadInventory()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	item, response := service.FindInventoryItem(items, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, item, "Get Inventory Item.", 200)
}

func PutInventoryItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadInventory()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	newItem, err := dal.ContentToInventoryItem(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	_, response := service.UpdateInventoryItem(items, newItem, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
	}
	config.SendResponse(w, r, newItem, "Put Inventory.", 200)
}

func DeleteInventoryItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadInventory()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.RemoveInventoryItem(items, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, nil, "Delete Inventory.", 200)
}
