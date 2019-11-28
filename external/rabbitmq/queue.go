/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package rabbitmq

import "github.com/streadway/amqp"

type Queue struct {
	Queue *amqp.Queue
	Name string
}