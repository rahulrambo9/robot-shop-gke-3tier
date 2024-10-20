package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "log"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "35.237.164.110:9092",
        "group.id":          "my-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }
    defer c.Close()

    topic := "rambo-topic"
    c.SubscribeTopics([]string{topic}, nil)

    log.Println("Waiting for messages...")

    for {
        msg, err := c.ReadMessage(-1)
        if err == nil {
            log.Printf("Consumed message: %s", string(msg.Value))
        } else {
            log.Printf("Error while consuming: %v", err)
        }
    }
}
