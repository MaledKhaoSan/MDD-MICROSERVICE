package producer

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	// import package config
	"github.com/MD-PROJECT/PRODUCT-SERVICE-WRITE-MODEL/internal/domain/model"
)

func PublishProductCreatedEvent(product model.Product) error {
	log.Printf("📦 Publishing PrderCreatedEvent: %+v\n", product)

	payload, err := json.Marshal(product)
	if err != nil {
		log.Printf("❌ Error marshaling product: %v", err)
		return err
	}

	resp, err := http.Post(
		"http://localhost:8081/producer/productService/publishProductCreatedEvent",
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
