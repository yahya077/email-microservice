package service

import (
	"bytes"
	"fmt"
	"github.com/yahya077/email-microservice/models"
	"html/template"
	"net/smtp"
	"os"
)

func SendEmail(payload models.EmailPayload) {
	from := "no-reply@yahyahindioglu.com"
	to := payload.To

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	smtpHost := os.Getenv("HOST")
	smtpPort := os.Getenv("PORT")

	auth := smtp.PlainAuth("", username, password, smtpHost)

	t, _ := template.ParseFiles("template/template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", payload.Subject, mimeHeaders)))

	e := t.Execute(&body, struct {
		Subject string
		Body    string
	}{
		Subject: payload.Subject,
		Body:    payload.Message,
	})

	if e != nil {
		fmt.Println(e)
		return
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
