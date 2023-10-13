package main

import (
	"github.com/streadway/amqp"
	"golang.org/x/time/rate"
	"log"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	limiter := rate.NewLimiter(10, 1) // Limit to 10 requests per second.

	msgs, _ := ch.Consume(
		"somequeue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for d := range msgs {
		if limiter.Allow() {
			// Message Processing
			log.Printf("Received a message: %s", d.Body)
		} else {
			// Processing too many messages
		}
	}
}
