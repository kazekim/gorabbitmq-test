/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"github.com/kazekim/gorabbitmq-test/external/rabbitmq"
)



func main() {


	c, err := rabbitmq.NewConnection()
	if err != nil {
		panic(err)
	}

	q, err := c.QueueDeclare("Kazekim")
	if err != nil {
		panic(err)
	}
	err = c.Publish(q, "Hello, Jirawat.Kim")
	if err != nil {
		panic(err)
	}

}