package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Define RabbitMQ server URL
	amqpServerURL := os.Getenv("RABBITMQ_DSN")

	// Connect to RabbitMQ server
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare the exchange
	exchangeName := os.Getenv("RABBITMQ_EXCHANGE")
	err = ch.ExchangeDeclare(
		exchangeName, // exchange name
		"direct",     // exchange type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	// Declare a queue
	q, err := ch.QueueDeclare(
		"",    // random queue name
		false, // non-durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		q.Name,                            // queue name
		os.Getenv("RABBITMQ_ROUTING_KEY"), // routing key
		exchangeName,                      // exchange
		false,                             // no-wait
		nil,                               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %s", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Create a channel to receive messages
	forever := make(chan bool)

	// Receive messages
	go func() {
		for d := range msgs {
			log.Println("Received a message:")

			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, d.Body, "", "  ")
			if err != nil {
				fmt.Println("Error beautifying JSON:", err)
				return
			}

			fmt.Println(prettyJSON.String())
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
