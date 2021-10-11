package main

import (
	"log"
	"os"
	"rabbitmq/utils"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := bodyFrom(os.Args)
	// 将消息发送到延时队列上
	err = ch.Publish(
		"",           // exchange 这里为空则不选择 exchange
		"test_delay", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  "5000", // 设置五秒的过期时间
		})
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
