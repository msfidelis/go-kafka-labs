package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	topic := os.Getenv("KAFKA_TOPIC")
	boostrap_servers := os.Getenv("KAFKA_BROKER")
	consumer_group := os.Getenv("CONSUMER_GROUP")

	reader := getReader(boostrap_servers, topic, consumer_group)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message from consumer group %s at topic/partition/offset %v/%v/%v: %s = %s\n", consumer_group, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

}

func getReader(bootstrap_servers, topic string, consumer_group string) *kafka.Reader {

	dialer := &kafka.Dialer{
		Timeout: 5 * time.Second,
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(bootstrap_servers, ","),
		Topic:   topic,
		GroupID: consumer_group,
		Dialer:  dialer,
	})

}
