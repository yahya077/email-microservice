package service

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

/*func SendEmail(msg string, receivers ...string) {
	from := "no-reply@yahyahindioglu.com"
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	auth := smtp.PlainAuth("", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), host)

	bodyMessage := []byte(fmt.Sprintf("The message that comes from kafka: %s", msg))

	e := smtp.SendMail(host+":"+port, auth, from, receivers, bodyMessage)

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("Email Sent Successfully: %s", msg)
}*/

func SendEmail(msg string, receivers ...string) {
	m := gomail.NewMessage()

	m.SetHeader("From", "no-reply@yahyahindioglu.com")

	m.SetHeader("To", receivers...)

	m.SetHeader("Subject", "No-Reply E-Mail")

	m.SetBody("text/plain", msg)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	d := gomail.NewDialer(os.Getenv("HOST"), port, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
