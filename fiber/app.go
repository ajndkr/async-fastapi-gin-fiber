package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		requestID := uuid.New()
		c.Locals("requestID", requestID.String())

		log.Printf("New request - %s\n", requestID)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(10 * time.Second)
		return c.JSON(fiber.Map{"message": "success"})
	})

	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
