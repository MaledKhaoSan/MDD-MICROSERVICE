package router

import (
	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/query"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes กำหนดเส้นทาง API และรับ `db *gorm.DB`
func SetupRoutes(app *fiber.App, db *gorm.DB) {

	// Read
	app.Get("/inventory/inv/:store_id", func(c *fiber.Ctx) error {
		return query.GetInventoryListByStoreEndpoint(c, db)
	})

	app.Get("/inventory/product/:product_id", func(c *fiber.Ctx) error {
		return query.GetProductDetailsEndpoint(c, db)
	})

}
