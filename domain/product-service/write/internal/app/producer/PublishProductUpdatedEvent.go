package producer

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-WRITE-MODEL/internal/domain/model"
)

// PublishProductUpdatedEvent - ส่ง Event ไป Kafka
func PublishProductUpdatedEvent(productUpdate model.Product) error {
	log.Printf("📦 Publishing ProductUpdatedStatusEvent: %+v", productUpdate)

	payload, err := json.Marshal(productUpdate)
	if err != nil {
		log.Printf("❌ Error marshaling order update: %v", err)
		return err
	}

	log.Printf("📤 Sending Payload to Publisher: %s", string(payload))

	resp, err := http.Post(
		"http://localhost:8081/producer/productService/publishProductUpdatedEvent",
		"application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		log.Printf("❌ Error calling Publisher Service: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("❌ Publisher Service returned status %d", resp.StatusCode)
		return err
	}
	return nil
}
