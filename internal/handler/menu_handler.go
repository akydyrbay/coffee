package handler

import (
	"net/http"

	"frappuccino/config"
	"frappuccino/internal/dal"
	"frappuccino/internal/service"
)

func PostMenuItem(w http.ResponseWriter, r *http.Request) {
	items, err := dal.ReadMenu()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	newItem, err := dal.ContentToMenu(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.AddMenuItem(items, newItem)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, newItem, "Post Menu Item", 201)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	items, err := dal.ReadMenu()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	config.SendResponse(w, r, items, "Get Menu", 200)
}

func GetMenuItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadMenu()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	item, response := service.FindMenuItem(items, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, item, "Get Menu Item", 200)
}

func PutMenuItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadMenu()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	newItem, err := dal.ContentToMenu(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	_, response := service.UpdateMenuItem(items, newItem, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
	}
	config.SendResponse(w, r, newItem, "Put Menu Item", 200)
}

func DeleteMenuItem(w http.ResponseWriter, r *http.Request, id string) {
	items, err := dal.ReadMenu()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.RemoveMenuItem(items, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, nil, "Delete Menu Item", 200)
}
