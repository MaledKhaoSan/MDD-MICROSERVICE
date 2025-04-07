// domain/notification-service/internal/utils/EmailSender.go
package utils

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

type EmailTemplateData struct {
	StoreName     string
	OrderID       string
	OrderPrice    string
	OrderQuantity string
	StoreEmail    string
}

func SendEmailJS(to string, data EmailTemplateData) error {
	from := "khxokhaokhxo000@gmail.com"
	password := "cioafzuhnbrrtqgo"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Load HTML template
	tmpl, err := template.ParseFiles("internal/utils/order_created_template.html")
	if err != nil {
		log.Printf("❌ Failed to parse template: %v", err)
		return err
	}

	// Render template with params
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Printf("❌ Failed to execute template: %v", err)
		return err
	}

	subject := "New Order Notification"
	message := []byte("Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" + body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Printf("❌ Failed to send email: %v", err)
		return err
	}

	log.Println("✅ Email sent successfully (HTML)")
	return nil
}
