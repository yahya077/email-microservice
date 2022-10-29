package main

import (
	"fmt"
	"github.com/yahya077/email-microservice/service"
)

type kafkaDummyConsumer struct {
	msg string
}

func (k *kafkaDummyConsumer) NewConsumer(msg string) {
	k.msg = msg
}

func (k *kafkaDummyConsumer) ReadMessage() string {
	return k.msg
}

func main() {
	kafkaConsumer := kafkaDummyConsumer{}

	kafkaConsumer.NewConsumer("Dummy email message body")

	msg := kafkaConsumer.ReadMessage()

	fmt.Println(msg)

	service.SendEmail(msg, "dummy-user@yahyahindioglu.com")
}
