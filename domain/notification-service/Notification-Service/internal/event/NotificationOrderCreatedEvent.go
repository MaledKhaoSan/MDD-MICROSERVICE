package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MD-PROJECT/NOTIFICATION-SERVICE/internal/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderCreatedStruct struct {
	OrderID       string  `json:"order_id" validate:"required"`
	StoreID       string  `json:"store_id" validate:"required"`
	OrderPrice    float64 `json:"order_price"`
	OrderQuantity int     `json:"order_quantity"`
}

type storeInformation struct {
	StoreID    string `gorm:"primaryKey;column:store_id"`
	StoreName  string `gorm:"column:store_name"`
	StoreEmail string `gorm:"column:store_email"`
}

func NotificationOrderCreatedEvent(db *gorm.DB, message []byte) error {
	log.Printf("✅ [Notification] Received OrderCreated event")

	var req OrderCreatedStruct
	validate := validator.New()

	if err := json.Unmarshal(message, &req); err != nil {
		log.Printf("❌ Error unmarshalling event: %v", err)
		return err
	}

	if err := validate.Struct(&req); err != nil {
		log.Printf("❌ Validation failed: %v", err)
		return err
	}

	var storeInfo storeInformation
	if err := db.First(&storeInfo, "store_id = ?", req.StoreID).Error; err != nil {
		log.Printf("❌ Failed to fetch store email: %v", err)
		return err
	}

	// ใช้ค่าจาก EmailJS Template
	emailParams := map[string]string{
		"to_email":    storeInfo.StoreEmail,
		"store_name":  storeInfo.StoreName,
		"order_id":    req.OrderID,
		"order_price": fmt.Sprintf("%.2f", req.OrderPrice),
		"order_qty":   fmt.Sprintf("%d", req.OrderQuantity),
	}

	// ส่งอีเมลผ่าน EmailJS
	if err := utils.SendEmailJS(storeInfo.StoreEmail, emailParams); err != nil {
		log.Printf("❌ Email sending failed: %v", err)
		return err
	}

	log.Printf("✅ Email sent successfully to %s", storeInfo.StoreEmail)
	return nil
}
