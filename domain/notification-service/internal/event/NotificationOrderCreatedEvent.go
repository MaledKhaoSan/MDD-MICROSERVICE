package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MD-PROJECT/NOTIFICATION-SERVICE/internal/model"
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

	var storeInfo model.Store_Information
	if err := db.First(&storeInfo, "store_id = ?", req.StoreID).Error; err != nil {
		log.Printf("❌ Failed to fetch store email: %v", err)
		return err
	}

	// ใช้ค่าจาก EmailJS Template
	emailData := utils.EmailTemplateData{
		StoreName:     storeInfo.StoreName,
		OrderID:       req.OrderID,
		OrderPrice:    fmt.Sprintf("%.2f", req.OrderPrice),
		OrderQuantity: fmt.Sprintf("%d", req.OrderQuantity),
		StoreEmail:    storeInfo.StoreEmail,
	}

	// ส่งอีเมลผ่าน EmailJS
	if err := utils.SendEmailJS(storeInfo.StoreEmail, emailData); err != nil {
		log.Printf("❌ Email sending failed: %v", err)
		return err
	}

	log.Printf("✅ Email sent successfully to %s", storeInfo.StoreEmail)
	return nil
}
