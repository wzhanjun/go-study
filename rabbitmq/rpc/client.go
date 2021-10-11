package main

import (
	"log"
	"math/rand"
	"os"
	"rabbitmq/utils"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	utils.FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}

func randomString(len int) string {
	byteSlice := make([]byte, len)
	for i := 0; i < len; i++ {
		byteSlice[i] = byte(randInt(65, 90))
	}
	return string(byteSlice)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn := utils.GetRabbitConn()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	corrId := randomString(32)

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	err = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       queue.Name,
			Body:          []byte(strconv.Itoa(n)),
		})

	utils.FailOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return

}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	utils.FailOnError(err, "Failed to convert arg to integer")
	return n
}
