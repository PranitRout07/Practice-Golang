package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a new Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter messages (press Ctrl+C to exit):")
	for scanner.Scan() {
		message := scanner.Text()
		if strings.ToLower(message) == "exit" {
			break
		}
		err := writer.WriteMessages(context.Background(),
			kafka.Message{Value: []byte(message)},
		)
		if err != nil {
			fmt.Printf("Error producing message: %v\n", err)
		}
	}
	writer.Close()
}
