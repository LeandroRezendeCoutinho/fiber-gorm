package main

import (
	"log"
	config "main/configs"

	"main/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${path} ${method} ${status} ${latency}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Sao_Paulo",
	}))

	config.Connect()

	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/:id", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/:id", handlers.UpdateDog)
	app.Delete("/dogs/:id", handlers.RemoveDog)

	log.Fatal(app.Listen(":3000"))
}
