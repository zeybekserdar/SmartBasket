package main

import (
	"SmartBasket/db"
	"SmartBasket/product"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main(){
	database,err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect to db %v",err)
	}

	repository := product.NewRepository(database)
	err = repository.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := product.NewService(repository)
	handler := product.NewHandler(service)

	app := fiber.New()
	app.Get("/products/:id",handler.Get)
	app.Post("/products/",handler.Create)
	app.Listen(":8000")

	fmt.Println(database.Config)
}