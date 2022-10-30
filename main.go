package main

import (
	"github.com/yahya077/email-microservice/services/emailerService"
	"github.com/yahya077/email-microservice/services/eventService"
)

func main() {
	eventService.Listen(emailerService.SendEmail)
}
