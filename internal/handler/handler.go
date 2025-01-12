package handler

import (
	"net/http"
	"regexp"
	"strings"

	. "frappuccino/config"
)

type RequestHandler struct{}

var (
	MenuRe           = regexp.MustCompile(`^\/menu\/?$`)
	MenuItemRe       = regexp.MustCompile(`^\/menu\/\w+\/?$`)
	InventoryRe      = regexp.MustCompile(`^\/inventory$`)
	InventoryItemRe  = regexp.MustCompile(`^\/inventory\/\w+\/?$`)
	OrdersRe         = regexp.MustCompile(`^\/orders\/?$`)
	OrderItemRe      = regexp.MustCompile(`^\/orders\/\w+\/?$`)
	OrderItemCloseRe = regexp.MustCompile(`^\/orders\/\w+/close\/?$`)
	PopularItemsRe   = regexp.MustCompile(`^\/reports\/popular-items\/?$`)
	TotalSalesRe     = regexp.MustCompile(`^\/reports\/total-sales\/?$`)
)

func (req *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && MenuRe.MatchString(r.URL.Path):
		GetMenu(w, r)
		return
	case r.Method == http.MethodPost && MenuRe.MatchString(r.URL.Path):
		PostMenuItem(w, r)
		return
	case r.Method == http.MethodGet && MenuItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		GetMenuItem(w, r, id)
		return
	case r.Method == http.MethodPut && MenuItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		PutMenuItem(w, r, id)
		return
	case r.Method == http.MethodDelete && MenuItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		DeleteMenuItem(w, r, id)
		return
	case r.Method == http.MethodPost && InventoryRe.MatchString(r.URL.Path):
		PostInventoryItem(w, r)
		return
	case r.Method == http.MethodGet && InventoryRe.MatchString(r.URL.Path):
		GetInventory(w, r)
		return
	case r.Method == http.MethodGet && InventoryItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		GetInventoryItem(w, r, id)
		return
	case r.Method == http.MethodPut && InventoryItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		PutInventoryItem(w, r, id)
		return
	case r.Method == http.MethodDelete && InventoryItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		DeleteInventoryItem(w, r, id)
		return
	case r.Method == http.MethodPost && OrdersRe.MatchString(r.URL.Path):
		PostOrder(w, r)
		return
	case r.Method == http.MethodGet && OrdersRe.MatchString(r.URL.Path):
		GetOrders(w, r)
		return
	case r.Method == http.MethodGet && OrderItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		GetOrder(w, r, id)
		return
	case r.Method == http.MethodPut && OrderItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		PutOrder(w, r, id)
		return
	case r.Method == http.MethodDelete && OrderItemRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		DeleteOrder(w, r, id)
		return
	case r.Method == http.MethodPost && OrderItemCloseRe.MatchString(r.URL.Path):
		id := strings.Split(r.URL.Path, "/")[2]
		CloseOrder(w, r, id)
		return
	case r.Method == http.MethodGet && PopularItemsRe.MatchString(r.URL.Path):
		GetPopularItems(w, r)
		return
	case r.Method == http.MethodGet && TotalSalesRe.MatchString(r.URL.Path):
		GetTotalSales(w, r)
		return
	default:
		Error(w, 400, "Invalid Request.")
		return
	}
}
