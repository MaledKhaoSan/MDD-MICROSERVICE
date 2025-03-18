package query

import (
	"log"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProductDetailsEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	productID := c.Params("product_id")

	// ✅ ค้นหาสินค้าพร้อมดึงข้อมูลหมวดหมู่
	var product model.Product
	if err := db.Preload("Category").
		Where("product_id = ?", productID).
		First(&product).Error; err != nil {

		// ✅ ถ้าไม่พบข้อมูล
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		// ✅ กรณีเกิด Error อื่น
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	log.Printf("✅ Get Product Details: \n%+v\n", product)
	return c.Status(fiber.StatusOK).JSON(product)
}
