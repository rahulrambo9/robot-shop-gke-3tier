package main

import (
    "log"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    topic := "rambo-topic" // Define your topic as a variable

    // Create a new producer instance
    producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "35.237.164.110:9092"})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer producer.Close()

    // Produce messages
    for _, word := range []string{"Hello", "Kafka", "from", "Go"} {
        // Create a new message
        message := &kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          []byte(word),
        }

        // Produce the message
        err := producer.Produce(message, nil)
        if err != nil {
            log.Printf("Failed to produce message: %s", err)
        }
    }

    // Wait for all messages to be delivered
    producer.Flush(15 * 1000)
}
