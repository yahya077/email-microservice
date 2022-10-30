package emailerService

import (
	"bytes"
	"fmt"
	"github.com/yahya077/email-microservice/models"
	"html/template"
	"net/smtp"
)

type HtmlEmail struct {
}

func (e HtmlEmail) Send(payload models.EmailPayload, email Email) {
	var body bytes.Buffer

	email.configure()
	t, _ := template.ParseFiles("template/template.html")

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", payload.Subject, mimeHeaders)))

	executeErr := t.Execute(&body, struct {
		Subject string
		Body    string
	}{
		Subject: payload.Subject,
		Body:    payload.Message,
	})

	if executeErr != nil {
		fmt.Println(e)
		return
	}

	sendErr := smtp.SendMail(email.smtpHost+":"+email.smtpPort, email.auth, payload.From, []string{payload.To}, body.Bytes())

	if sendErr != nil {
		fmt.Println(sendErr)
		return
	}

	fmt.Println("Email Sent!")
}
