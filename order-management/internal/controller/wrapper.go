package controller

import (
	"encoding/json"
	"my-ecommerce-shop/order-management/internal/contract"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) *contract.GenericResponse

func Wrap(handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := handler(w, r)

		if response != nil {
			resp, _ := json.Marshal(response)
			_, _ = w.Write(resp)
		}
	}
}
