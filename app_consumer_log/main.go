package main

import (
	"fmt"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	topic := os.Getenv("KAFKA_TOPIC")
	boostrap_servers := os.Getenv("KAFKA_BROKER")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     boostrap_servers,
		"broker.address.family": "v4",
		"group.id":              "oi",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

	fmt.Printf("Closing consumer\n")
	c.Close()
}
