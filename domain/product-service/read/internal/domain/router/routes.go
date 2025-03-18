package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-READ-MODEL/internal/app/command"
	"github.com/MD-PROJECT/PRODUCT-SERVICE-READ-MODEL/internal/app/query"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// CQRS: Command Handlers
	app.Put("/product/api", func(c *fiber.Ctx) error {
		return command.UpdateProductEndpoint(c, db)
	})

	// CQRS: Query Handlers
	app.Get("/product/api/:id", func(c *fiber.Ctx) error {
		return query.GetProductDetailEndpoint(c, db)
	})
}
