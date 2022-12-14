package main

import (
	"fmt"
	"helloworld/broker"
	"os"

	"github.com/pkg/errors"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch, err := broker.RabbitMQ()

	if err != nil {
		panic(err)
	}

	defer func() {
		ch.Close()
		conn.Close()
	}()

	q, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)

	if err != nil {
		panic(errors.Wrap(err, "failed to declare queue"))
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		Headers:     map[string]interface{}{},
		ContentType: "text/plain",
		Body:        []byte(os.Args[1]),
	})

	if err != nil {
		panic(errors.Wrap(err, "failed to publish message"))
	}

	fmt.Println("Send message:", os.Args[1])
}
