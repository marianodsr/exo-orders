package orders

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/marianodsr/exo-orders/storage"
	"gorm.io/gorm"
)

type status string

type Order struct {
	gorm.Model
	DeliveryDate time.Time `json:"delivery_date"`
	EmmisionDate time.Time `json:"emission_date"`
	Status       status    `sql:"type:status" json:"status"`
}

const (
	Cancelled status = "cancelled"
	Pending   status = "pending"
	Fulfilled status = "fulfilled"
)

func (p *status) Scan(value interface{}) error {
	*p = status(value.([]byte))
	return nil
}

func (p status) Value() (driver.Value, error) {
	return string(p), nil
}

func (Order) TableName() string {
	return "orders"
}

func createOrder(order *Order) error {
	db := storage.GetDbConnection()
	res := db.Create(order)
	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("error creating order")
	}
	return nil
}

func getOrders() (Order, error) {
	db := storage.GetDbConnection()
	var order Order
	db.First(&order)
	fmt.Println("HERE")

	return order, nil

}
