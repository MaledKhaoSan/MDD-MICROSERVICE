package infra

import (
	"log"

	order "github.com/MD-PROJECT/SUBSCRIBE-SERVICE/internal/app/order-service"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Define callback function type
type EventHandler func([]byte)

// Mapping: Topic -> Multiple Handlers
var topicHandlers = map[string][]EventHandler{
	"OrderCreatedEvent": {
		order.OrderUpdate,
		// inventory.InventoryDecrease,
		// notification.NotificationNewOrder,
	},
	"OrderUpdatedStatusEvent": {
		order.OrderUpdate,
		// dashboard.DashBoardOrderUpdateStatus,
	},
}

func StartKafkaConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"group.id":          "subscribe-service",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("❌ Kafka Consumer Error: %v", err)
	}
	defer consumer.Close()

	// Subscribe to all topics dynamically
	topics := []string{}
	for topic := range topicHandlers {
		topics = append(topics, topic)
	}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("❌ Failed to subscribe topics: %v", err)
	}

	log.Println("✅ Kafka Consumer started...")

	for {
		msg, err := consumer.ReadMessage(-1) // Blocking call
		if err == nil {
			log.Printf("📥 Received Event:\n\n %s from %s\n", string(msg.Value), *msg.TopicPartition.Topic)

			// Find the correct handlers for this topic
			if handlers, exists := topicHandlers[*msg.TopicPartition.Topic]; exists {
				for _, handler := range handlers {
					go handler(msg.Value) // Run asynchronously
				}
			} else {
				log.Printf("⚠️ No handler found for topic: %s", *msg.TopicPartition.Topic)
			}
		} else {
			log.Printf("❌ Kafka Error: %v\n", err)
		}
	}
}
