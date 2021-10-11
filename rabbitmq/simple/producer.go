package main

import (
	"log"
	"rabbitmq/utils"

	amqp "github.com/streadway/amqp"
)

func main() {

	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "failed to declare a queue")

	body := "hello world!"
	err = ch.Publish(
		"",         // echange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	utils.FailOnError(err, "failed to publish a message")
	log.Printf("[x] Send %s", body)
}
