package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/ORDER-HISTORY-SERVICE/internal/query"
)

// SetupRoutes กำหนดเส้นทาง API และรับ `db *gorm.DB`
func SetupRoutes(app *fiber.App, db *gorm.DB) {

	// Read
	app.Get("/orders/:id", func(c *fiber.Ctx) error {
		return query.GetOrderDetailsEndpoint(c, db)
	})
}
