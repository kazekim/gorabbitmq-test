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

	q, err := c.ExchangeDeclareSenderBroadcast("KazekimBC")
	if err != nil {
		panic(err)
	}
	message := "Hello, Jirawat.Kim"
	err = c.PublishBroadcast(q, message)
	if err != nil {
		panic(err)
	}
	fmt.Println("Send Message: Hello, Jirawat.Kim")
}
