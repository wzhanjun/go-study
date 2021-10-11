package main

import (
	"log"
	"os"
	"rabbitmq/utils"
	"strings"

	"github.com/streadway/amqp"
)

// 1. 配置连接工厂
// 2. 建立TCP连接
// 3. 在TCP连接的基础上创建通道
// 4. 声明一个topic交换机
// 5. 发送消息，并配置消息的路由键

func main() {
	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		// "",    // name
		severityFrom(os.Args), // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name,   // queue name
		queue.Name,   // routing key
		"logs_topic", // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic", // exchange
		// severityFrom(os.Args), // routing key
		queue.Name,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
