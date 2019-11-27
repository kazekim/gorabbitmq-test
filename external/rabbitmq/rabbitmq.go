/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	RabbitMQConnectionString = "amqp://guest:guest@localhost:15672/"
)

type Connection struct {
	conn *amqp.Connection
	channel *amqp.Channel
}

func NewConnection() (*Connection, error) {
	conn, err := amqp.Dial(RabbitMQConnectionString)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err.Error())
	}

	return &Connection{
		conn: conn,
		channel: ch,
	}, nil
}

func (c *Connection) QueueDeclare(name string) (*Queue, error) {
	q, err := c.channel.QueueDeclare(
		name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err.Error())
	}

	return &Queue{
		q: q,
	}, nil
}

func (c *Connection) Close() {
	c.conn.Close()
	c.channel.Close()
}

func (c *Connection) Publish(q *Queue, message string) error {
	body := "Hello World!"
	err := c.channel.Publish(
		"",     // exchange
		q.q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err.Error())
	}
	return nil
}

func (c *Connection) Consume(q *Queue, f func(message string)) error {
	msgs, err := c.channel.Consume(
		q.q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			f(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}