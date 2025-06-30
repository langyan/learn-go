package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Middleware to measure time
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		elapsed := time.Since(start)
		log.Printf("Fiber: %s %s took %s\n", c.Method(), c.Path(), elapsed)
		return err
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	log.Fatal(app.Listen(":8080"))
}
