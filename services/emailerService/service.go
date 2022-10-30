package emailerService

import (
	"github.com/yahya077/email-microservice/models"
	"net/smtp"
	"os"
)

var EmailInit Email

func init() {
	EmailInit.configure()
}

type IEmailer interface {
	Send(payload models.EmailPayload, email Email)
}

type Email struct {
	username string
	password string
	smtpHost string
	smtpPort string
	auth     smtp.Auth
	IEmailer IEmailer
}

func (e *Email) configure() {
	e.username = os.Getenv("USERNAME")
	e.password = os.Getenv("PASSWORD")
	e.smtpHost = os.Getenv("HOST")
	e.smtpPort = os.Getenv("PORT")
	e.setPlainAuth()
}

func (e *Email) setPlainAuth() {
	e.auth = smtp.PlainAuth("", e.username, e.password, e.smtpHost)
}

func LoadEmailer(payload models.EmailPayload) IEmailer {
	switch payload.Type {
	case "text":
		//TODO: change with TextEmail
		return HtmlEmail{}
	default:
		return HtmlEmail{}
	}

}

func Send(payload models.EmailPayload, emailer IEmailer) {
	emailer.Send(payload, EmailInit)
}
