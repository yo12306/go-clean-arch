package main

import (
	"log"

	"github.com/go-clean-arch/adapters"
	"github.com/go-clean-arch/entities"
	"github.com/go-clean-arch/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(&entities.Order{}); err != nil {
		log.Fatal("failed to migrate database: %v", err)
	}

	orderRepo := adapters.NewGormRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8000")
}
