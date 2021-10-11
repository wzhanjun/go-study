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
// 4. 声明一个fanout交换机 (而不是声明一个队列)
// 5. 发送消息
func main() {
	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
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
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
