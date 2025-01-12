package handler

import (
	"net/http"

	"frappuccino/config"
	"frappuccino/internal/dal"
	"frappuccino/internal/service"
)

func GetPopularItems(w http.ResponseWriter, r *http.Request) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	sortedItems, response := service.SortByPopularity(orders)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	config.SendResponse(w, r, sortedItems, "Get Popular Items.", 200)
}

func GetTotalSales(w http.ResponseWriter, r *http.Request) {
	orders, err := dal.ReadOrders()
	if err != nil {
		config.Error(w, 500, err.Error())
		return
	}
	totalSales, response := service.SumSales(orders)
	if response.ErrorMessage != nil {
		config.Error(w, response.Code, response.ErrorMessage.Error())
		return
	}
	// map[string]string{"error": text}
	config.SendResponse(w, r, map[string]float64{"total_sales": totalSales}, "Get Total Sales.", 200)
}
