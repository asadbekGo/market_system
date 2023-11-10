package main

import (
	"fmt"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/controller"
	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/storage/postgres"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, pgStorage)

	category(con)
}

func category(con *controller.Conn) {

	// for i := 0; i < 100; i++ {
	// 	con.CreateCategory(&models.CreateCategory{
	// 		Title:    faker.FirstName(),
	// 		ParentID: "689dfd45-8166-402e-a490-d044d843694f",
	// 	})
	// }

	// resp, err := con.GetListCategory(&models.GetListCategoryRequest{
	// 	Offset: 0,
	// 	Limit:  100,
	// 	Search: "m",
	// })

	// if err != nil {
	// 	log.Println("Error while GetListCategory >>> " + err.Error())
	// 	return
	// }

	// fmt.Println("Category Count:", resp.Count)
	// for _, category := range resp.Categories {
	// 	fmt.Println(category.Title)
	// }

	resp, err := con.UpdateCategory(&models.UpdateCategory{
		Id:       "64c5ebb8-6b43-406e-b285-a3009f9cf3e9",
		Title:    "JUBAJUBA",
		ParentID: "",
	})

	if err != nil {
		log.Println("Error while UpdateCategory >>> " + err.Error())
		return
	}

	fmt.Println(resp)
}
