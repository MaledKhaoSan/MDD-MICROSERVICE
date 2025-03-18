package query

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-READ-MODEL/internal/domain/model"
)

// GetProductDetailEndpoint ดึงข้อมูล product ตาม product_id (Sync)
func GetProductDetailEndpoint(c *fiber.Ctx, db *gorm.DB) error {

	productID := c.Params("id")

	// ค้นหา Product ใน Database
	var product model.Product
	if err := db.Where("product = ?", productID).First(&product).Error; err != nil {
		// ถ้าไม่พบข้อมูล
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		// กรณีเกิด Error อื่น
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	log.Printf("✅ Get Product Detail: \n%+v\n", product)
	return c.Status(fiber.StatusOK).JSON(product)
}
