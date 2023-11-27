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
	app.Get("/entities",returnEntities)
	newApi := app.Group("/newapi")
	newApi.Get("/",api)
	app.Listen(":8080")

}
func returnEntities(c *fiber.Ctx) error{
	return c.JSON(JSONTextResponse{Message: "Message from entity path!!"})
	// return c.SendString("Message from entity path!!")
}
func api(c *fiber.Ctx) error{
	return c.JSON(JSONTextResponse{Message: "Message from api path!!"})
}