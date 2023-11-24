package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)
type JSONTextResponse struct{
	Message string
}
func main() {
	fmt.Println("hello")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello!!!")
		return c.JSON(JSONTextResponse{Message: "Hello,whats up"})
	})
	app.Listen(":8080")

}
