package main

import (
	"github.com/marianodsr/exo-orders/router"
	"github.com/marianodsr/exo-orders/storage"
)

func main() {
	storage.InitDB()

	router.HandleRoutes()

}
