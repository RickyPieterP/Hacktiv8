package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Orders struct {
	OrderID      string    `json:"orderID" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items" example:"1"`
}

type Item struct {
	ItemID      string `json:"itemID" example:"ABC123"`
	Description string `json:"description" example:"A Random Description"`
	Quantity    int    `json:"quantity" example:"1"`
}

// @title Orders API
// @version 1.0
// @desciption This is a sample service of managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200
// Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
