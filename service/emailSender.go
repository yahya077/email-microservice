package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendEmail(msg string, receivers ...string) {
	from := "no-reply@yahyahindioglu.com"
	to := receivers

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	smtpHost := os.Getenv("HOST")
	smtpPort := os.Getenv("PORT")

	auth := smtp.PlainAuth("", username, password, smtpHost)

	t, _ := template.ParseFiles("template/template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	e := t.Execute(&body, struct {
		Subject string
		Body    string
	}{
		Subject: "This is a test subject",
		Body:    msg,
	})

	if e != nil {
		fmt.Println(e)
		return
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
