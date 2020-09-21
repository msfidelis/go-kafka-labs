package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	guuid "github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func main() {

	topic := os.Getenv("KAFKA_TOPIC")
	boostrap_servers := os.Getenv("KAFKA_BROKER")

	writer := getWriter(boostrap_servers, topic)
	defer writer.Close()

	// Async Producer
	// for {
	// 	r := make(chan int32)

	// 	// Async Producer
	// 	go func() {

	// 		msg := kafka.Message{
	// 			Key:   []byte(guuid.New().String()),
	// 			Value: []byte(guuid.New().String()),
	// 		}

	// 		err := writer.WriteMessages(context.Background(), msg)

	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		r <- 1
	// 	}()
	// }

	// Sync Producer
	for {

		msg := kafka.Message{
			Key:   []byte(guuid.New().String()),
			Value: []byte(guuid.New().String()),
		}

		err := writer.WriteMessages(context.Background(), msg)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second)

	}

}

func getWriter(bootstrap_servers, topic string) *kafka.Writer {

	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: name,
	}

	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(bootstrap_servers, ","),
		Topic:   topic,
		// Balancer:     &kafka.LeastBytes{},
		Balancer:     &kafka.Hash{},
		Dialer:       dialer,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	})

}
