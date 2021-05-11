package orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router) {

	r.Post("/", createOrderAttempt)
	r.Get("/", getAllOrders)

}

func createOrderAttempt(w http.ResponseWriter, r *http.Request) {
	order := &Order{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid body format", http.StatusBadRequest)
		return
	}
	// HERE WE WOULD EMIT AN ORDER EMITTED EVENT WHICH CATALOG MICROSERVICE WOULD LISTEN FOR AND APPROVE IF
	// SUFICIENT STOCK
	if err := createOrder(order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := getOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "an error ocurred", http.StatusInternalServerError)
		return
	}
}
