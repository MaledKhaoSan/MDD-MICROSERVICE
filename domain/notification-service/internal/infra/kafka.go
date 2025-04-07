package infra

import (
	"log"

	"github.com/MD-PROJECT/NOTIFICATION-SERVICE/internal/event"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

// Define callback function type with db parameter
type EventHandler func(*gorm.DB, []byte) error

// Modified StartKafkaConsumer to accept db parameter
func StartKafkaConsumer(db *gorm.DB) { // Added db parameter
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"group.id":          "notification-subscribe-service-group",
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
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("📥 Received Event:\n\n %s from %s\n", string(msg.Value), *msg.TopicPartition.Topic)

			if handlers, exists := topicHandlers[*msg.TopicPartition.Topic]; exists {
				for _, handler := range handlers {
					go func(h EventHandler) { // Wrap in closure to capture handler
						if err := h(db, msg.Value); err != nil {
							log.Printf("❌ Handler error: %v", err)
						}
					}(handler)
				}
			} else {
				log.Printf("⚠️ No handler found for topic: %s", *msg.TopicPartition.Topic)
			}
		} else {
			log.Printf("❌ Kafka Error: %v\n", err)
		}
	}
}

// Mapping: Topic -> Multiple Handlers
var topicHandlers = map[string][]EventHandler{
	"OrderCreatedEvent": {
		event.NotificationOrderCreatedEvent,
	},
}
