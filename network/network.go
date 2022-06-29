package network

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamidreza01/blueberry/config"
)

func Start(conf config.Config) {
	app := fiber.New(fiber.Config{})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Listen(conf.Ip + conf.Port)
}
