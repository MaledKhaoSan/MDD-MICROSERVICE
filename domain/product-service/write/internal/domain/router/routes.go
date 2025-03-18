package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-WRITE-MODEL/internal/app/command"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// CQRS: Command Handlers
	app.Post("/product/api", func(c *fiber.Ctx) error {
		return command.CreateProductEndpoint(c, db)
	})
	app.Put("/product/api", func(c *fiber.Ctx) error {
		return command.UpdateProductEndpoint(c, db)
	})
}
