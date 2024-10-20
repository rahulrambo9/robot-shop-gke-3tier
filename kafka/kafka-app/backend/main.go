package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "sync"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    _ "github.com/lib/pq"
)

var (
    producer     *kafka.Producer
    consumer     *kafka.Consumer
    db           *sql.DB
    mu           sync.Mutex
    messages     []string
    topicName    = "rambo-topic" // Declare a variable for the topic name
)

func init() {
    // Kafka Producer
    var err error
    producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "35.237.164.110:9092"})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }

    // Kafka Consumer
    consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "35.237.164.110:9092",
        "group.id":          "my-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }
    consumer.SubscribeTopics([]string{topicName}, nil) // Use the variable here

    // PostgreSQL DB
    db, err = sql.Open("postgres", "user=postgres password=UTtf7mWwiy dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %s", err)
    }
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Message string `json:"message"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    // Produce the message
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny}, // Use the variable here
        Value:          []byte(req.Message),
    }
    producer.Produce(message, nil)

    // Store the message in the database
    mu.Lock()
    messages = append(messages, req.Message)
    mu.Unlock()

    w.WriteHeader(http.StatusOK)
}

func getMessages(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    json.NewEncoder(w).Encode(messages)
}

func consumeMessages() {
    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
            log.Printf("Consumed message: %s", string(msg.Value))
            // Store in database
            mu.Lock()
            messages = append(messages, string(msg.Value))
            mu.Unlock()
        }
    }
}

func main() {
    go consumeMessages()

    http.HandleFunc("/send", sendMessage)
    http.HandleFunc("/messages", getMessages)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
