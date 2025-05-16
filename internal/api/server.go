package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadduli/ShipSwift/config"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)
	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "I am breathing!",
	})
}
