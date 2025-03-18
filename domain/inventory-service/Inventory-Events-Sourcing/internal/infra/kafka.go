package infra

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

// InitKafkaProducer สร้างและกำหนดค่า Kafka producer
func InitKafkaProducer() {
	brokers := []string{"localhost:9093"} // ใช้ port 9093 ตาม EXTERNAL listener

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 100 * time.Millisecond

	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to initialize Sarama producer: %v", err)
	}
}

// PublishToKafka ส่งข้อความไปยัง Kafka topic
func PublishToKafka(topic string, payload []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to topic %s: %v", topic, err)
		return err
	}
	log.Printf("Message sent to topic %s, partition %d, offset %d", topic, partition, offset)
	return nil
}

// CloseKafkaProducer ปิด producer เมื่อไม่ใช้งาน
func CloseKafkaProducer() {
	if err := producer.Close(); err != nil {
		log.Printf("Failed to close Sarama producer: %v", err)
	}
}
