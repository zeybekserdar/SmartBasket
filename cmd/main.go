package main

import (
	"SmartBasket/db"
	"SmartBasket/product"
	"SmartBasket/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	productRepository := product.NewRepository(database)
	userRepository := user.NewRepository(database)
	err = productRepository.Migration()
	err = userRepository.Migration()
	if err != nil {
		log.Fatal(err)
	}
	productService := product.NewService(productRepository)
	userService := user.NewService(userRepository)
	productHandler := product.NewHandler(productService)

	userHandler := user.NewUserHandler(userService)

	app := fiber.New()
	app.Use(logger.New())

	//app.Use(requestid.New())

	app.Get("api/products/:id", productHandler.Get)
	app.Get("api/products", productHandler.GetAll)
	app.Delete("api/products/:id", productHandler.Delete)
	app.Put("api/products/:id", productHandler.Update)
	app.Post("api/products/", productHandler.Create)

	app.Post("api/users/login", userHandler.Login)
	app.Post("api/users/register", userHandler.Register)

	// Monitor
	app.Get("/monitor", monitor.New())

	appErr := app.Listen(":8000")
	if appErr != nil {
		log.Fatalf("Application failed", appErr)
	}

	fmt.Println(database.Config)
}
