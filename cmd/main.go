package main

import (
	"log"
	"net/http"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/controller"
	"github.com/asadbekGo/market_system/storage/postgres"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	handler := controller.NewController(&cfg, pgStorage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)

	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	if err := http.ListenAndServe(cfg.ServiceHost+cfg.ServiceHTTPPort, nil); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
