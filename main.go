package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/yahya077/email-microservice/models"
	"github.com/yahya077/email-microservice/service"
	"os"
)

var emailPayload models.EmailPayload

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVER"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":    os.Getenv("SASL_MECHANISM"),
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	consumer.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC")}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return
		}

		json.Unmarshal(msg.Value, &emailPayload)

		fmt.Printf("Message on %s: %s\n", emailPayload.To, emailPayload.To)

		service.SendEmail(emailPayload)

	}

	consumer.Close()
}
