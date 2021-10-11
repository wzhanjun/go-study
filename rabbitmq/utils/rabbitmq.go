package utils

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	rabbitConn *amqp.Connection
)

func GetRabbitConn() *amqp.Connection {
	if rabbitConn != nil {
		return rabbitConn
	}
	rabbitConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	FailOnError(err, "failed to connect to rabbitmq")

	return rabbitConn
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
