package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/transactrx/testFiber/pkg/appHelper"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app2 := appHelper.App{
		Name:    "TestApp",
		Version: "V1",
	}
	app2.GetName()

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
