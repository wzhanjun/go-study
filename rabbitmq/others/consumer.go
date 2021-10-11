package main

import (
	"log"
	"os"
	"rabbitmq/utils"
)

// 1. 配置连接工厂
// 2. 建立TCP连接
// 3. 在TCP连接的基础上创建通道
// 4. 声明一个topic交换机
// 5. 声明一个临时队列
// 6. 将临时队列绑定到交换机上，并在队列上绑定多个绑定值
// 7. 接收消息并处理
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
		"order.refund", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s", queue.Name, "logs_direct", s)
		err = ch.QueueBind(
			queue.Name,   // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
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
