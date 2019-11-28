/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (c *Connection) ExchangeDeclareSenderBroadcast(name string) (*Queue, error) {
	err := c.channel.ExchangeDeclare(
		name,   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare an exchange: %v", err.Error())
	}

	q := &Queue{
		Queue: nil,
		Name: name,
	}

	return q, nil
}

func (c *Connection) Publish(q *Queue, message string) error {
	err := c.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err.Error())
	}
	return nil
}

func (c *Connection) PublishPersistent(q *Queue, message string) error {
	err := c.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			//we're sure that the task_queue queue won't be lost even if RabbitMQ restarts.
			// Now we need to mark our messages as persistent - by using the amqp.Persistent option
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err.Error())
	}
	return nil
}

func (c *Connection) PublishBroadcast(q *Queue, message string) error {
	err := c.channel.Publish(
		q.Name,     // exchange
		"", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err.Error())
	}
	return nil
}