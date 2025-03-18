package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-EVENTS-SOURCING/internal/command"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/product", func(c *fiber.Ctx) error {
		return command.CreateProductEndpoint(c, db)
	})

	app.Patch("/product/:product_id", func(c *fiber.Ctx) error {
		return command.UpdateProductEndpoint(c, db)
	})

}
