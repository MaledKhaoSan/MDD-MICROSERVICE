package utils

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"time"
)

// RetryHTTP ส่ง HTTP Request พร้อม Retry Mechanism
func RetryHTTP(method, url string, payload []byte, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := SendHTTPRequest(method, url, payload)
		if err == nil {
			return nil
		}
		log.Printf("🔄 Retrying (%d/%d) for URL: %s", i+1, maxRetries, url)
		time.Sleep(2 * time.Second) // รอ 2 วินาทีก่อน retry
	}
	return errors.New("❌ Max retries reached for: " + url)
}

// SendHTTPRequest ใช้ส่ง HTTP Request (รองรับ GET, POST, PUT, PATCH)
func SendHTTPRequest(method, url string, payload []byte) error {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Failed to reach service: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("❌ Service returned status: %d", resp.StatusCode)
		return errors.New("Service returned unexpected status")
	}

	log.Printf("✅ Service returned status: %d", resp.StatusCode)
	return nil
}
