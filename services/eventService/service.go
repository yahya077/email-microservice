package eventService

import (
	"encoding/json"
	"fmt"
	"github.com/yahya077/email-microservice/models"
	"github.com/yahya077/email-microservice/services/emailerService"
	"os"
)

func Listen(callBackFunc func(payload models.EmailPayload, emailer emailerService.IEmailer)) {
	var emailPayload models.EmailPayload

	consumer.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC")}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return
		}

		json.Unmarshal(msg.Value, &emailPayload)

		callBackFunc(emailPayload, emailerService.LoadEmailer(emailPayload))
	}

	consumer.Close()
}
