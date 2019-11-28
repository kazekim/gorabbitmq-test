/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package rabbitmq

import (
	"fmt"
	"log"
	)

func (c *Connection) EnableQOS(prefetchCount, prefetchSize int, global bool) error {
	err := c.channel.Qos(
		prefetchCount,     // prefetch count (Recommend : 1)
		prefetchSize,     // prefetch size (Recommend : 0)
		false, // global
	)
	return err
}

func (c *Connection) ExchangeDeclareReceiverBroadcast(name string) (*Queue, error) {
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

	queue, err := c.channel.QueueDeclare(
		"", // name
		//When RabbitMQ quits or crashes it will forget the queues and messages unless you tell it not to.
		// Two things are required to make sure that messages aren't lost: we need to mark both the queue and messages as durable.
		false,   // durable
		false,   // delete when unused
		true,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err.Error())
	}

	q := Queue{
		Queue: &queue,
		Name: "",
	}

	err = c.channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		name, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %v", err.Error())
	}

	return &q, nil
}

func (c *Connection) Consume(q *Queue, autoAck bool, f func(message string)) error {
	msgs, err := c.channel.Consume(
		q.Name, // queue
		"",     // consumer
		//Using auto-ack = false we can be sure that even if you kill a worker using CTRL+C while it was processing a message,
		// nothing will be lost. Soon after the worker dies all unacknowledged messages will be redelivered.
		//Acknowledgement must be sent on the same channel that received the delivery. Attempts to acknowledge using
		// a different channel will result in a channel-level protocol exception
		autoAck,   // auto-ack
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

