package handler

import (
	"net/http"

	"frappuccino/config"
	"frappuccino/internal/dal"
	"frappuccino/internal/service"
)

// NOTE: Handler is about how data looks not how it operates
// TODO: fix naming
// TODO: Send "ok" response for success operations if it needs

func PostOrder(w http.ResponseWriter, r *http.Request) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	newOrder, err := dal.ContentToOrder(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	response := service.AddOrder(orders, newOrder)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, newOrder, "Post Order.", 201)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	config.SendResponse(w, r, orders, "Get Orders.", 200)
}

func GetOrder(w http.ResponseWriter, r *http.Request, id string) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	order, response := service.FindOrder(orders, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, order, "Get Order.", 200)
}

func PutOrder(w http.ResponseWriter, r *http.Request, id string) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	newOrder, err := dal.ContentToOrder(r.Body)
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.UpdateOrder(orders, newOrder, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
	}
	config.SendResponse(w, r, nil, "Put Order.", 200)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}

	response := service.RemoveOrder(orders, id)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, nil, "Delete Order.", 200)
}

func CloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	response := service.FinishOrder(orders, id, "closed")
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, nil, "Close Order.", 200)
}
