package healthhandler

import (
	"github.com/gofiber/fiber/v2"

	healthmodel "go-rest-api-boilerplate/internal/healthhandler/model"
)

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(
		healthmodel.HealthEndpoint{
			Status: "OK",
		},
	)
}
