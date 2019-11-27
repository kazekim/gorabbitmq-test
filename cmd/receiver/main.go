/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/gorabbitmq-test/external/rabbitmq"
)

func main() {

	c, err := rabbitmq.NewConnection()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	q, err := c.QueueDeclare("Kazekim")
	if err != nil {
		panic(err)
	}

	fmt.Println("======Start Receiver======")

	f := func(message string) {
		fmt.Println("Receive : ", message)
	}

	c.Consume(q, f)
}