package main

import (
	"log"
	"rabbitmq/utils"

	"github.com/streadway/amqp"
)

func main() {
	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"test",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		"test_logs", // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name, // queue name, 这里指的是 test_logs
		"",         // routing key
		"test",     // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")

	// 申明一个延时队列
	_, errDelay := ch.QueueDeclare(
		"test_delay", // name
		false,        // durable
		false,        // delete when unused
		true,         // exclusive
		false,        // no-wait
		amqp.Table{
			// 当消息过期时把消息发送到 test 这个 exchange
			"x-dead-letter-exchange": "test",
		}, // arguments
	)
	utils.FailOnError(errDelay, "Failed to declare a delay_queue")

	// 这里监听的是 test_logs
	msgs, err := ch.Consume(
		queue.Name, // queue name, 这里指的是 test_logs
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
