package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer ...
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	producer, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}

	return producer
}

// Publish ...
func Publish(message string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	producedMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(message),
	}

	err := producer.Produce(producedMessage, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

// DeliveryReport ...
func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed: ", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to: ", ev.TopicPartition)
			}
		}
	}
}
