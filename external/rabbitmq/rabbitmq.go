/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
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
		//When RabbitMQ quits or crashes it will forget the queues and messages unless you tell it not to.
		// Two things are required to make sure that messages aren't lost: we need to mark both the queue and messages as durable.
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
		Queue: &q,
		Name: name,
	}, nil
}

func (c *Connection) Close() {
	c.conn.Close()
	c.channel.Close()
}