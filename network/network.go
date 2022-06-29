package network

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hamidreza01/blueberry/config"
)

func Start(conf config.Config) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if er, ok := err.(*fiber.Error); ok {
				switch er.Code {
				case 404:
					return c.SendFile(conf.Error404)
				case 500:
					return c.SendFile(conf.Error500)
				case 403:
					return c.SendFile(conf.Error403)
				case 203:
					return c.SendFile(conf.Error203)
				default:
					c.Set("Content-Type", "text/html")
					return c.SendString(RenderString(conf.Error, map[string]interface{}{
						"Error":   strconv.Itoa(er.Code),
						"Message": er.Message,
					}))
				}
			}
			c.Set("Content-Type", "text/html")
			return c.SendString(RenderString(conf.Error, map[string]interface{}{
				"Error":   "Unknown",
				"Message": "Unknown Error",
			}))
		},
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Listen(conf.Ip + conf.Port)
}
