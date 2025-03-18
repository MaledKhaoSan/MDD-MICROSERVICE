package query

import (
	"log"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetInventoryListByStoreEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	storeID := c.Params("store_id")

	// ✅ ค้นหา inventory ทั้งหมดของร้านค้าตาม `store_id`
	var inventoryList []model.Inventory
	if err := db.Preload("Product").Preload("Warehouse").
		Where("store_id = ?", storeID).
		Find(&inventoryList).Error; err != nil {

		// ✅ กรณี Error อื่น
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// ✅ ถ้าไม่มีสินค้าใน inventory
	if len(inventoryList) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No inventory found for this store"})
	}

	log.Printf("✅ Get Inventory List for Store ID %s: \n%+v\n", storeID, inventoryList)
	return c.Status(fiber.StatusOK).JSON(inventoryList)
}
