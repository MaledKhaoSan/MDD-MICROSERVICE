package utils

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// EmailJS API Endpoint
const emailJSEndpoint = "https://api.emailjs.com/api/v1.0/email/send"

// SendEmailJS ใช้ EmailJS API เพื่อส่งอีเมลโดยใช้ Template ที่ตั้งค่าไว้ใน EmailJS
func SendEmailJS(toEmail string, params map[string]string) error {
	// โหลด environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("❌ Error loading .env file: %v", err)
		return err
	}

	// ดึงค่าจาก .env
	serviceID := os.Getenv("EMAILJS_SERVICE_ID")
	templateID := os.Getenv("EMAILJS_TEMPLATE_ID")
	publicKey := os.Getenv("EMAILJS_PUBLIC_KEY")

	// JSON Payload สำหรับ EmailJS
	payload := map[string]interface{}{
		"service_id":      serviceID,
		"template_id":     templateID,
		"user_id":         publicKey,
		"template_params": params, // ใช้ params จาก template
	}

	// ส่ง HTTP POST ไปที่ EmailJS API
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(emailJSEndpoint)

	if err != nil {
		log.Printf("❌ Email sending failed: %v", err)
		return err
	}

	log.Printf("✅ Email sent successfully! Response: %s", resp)
	return nil
}
