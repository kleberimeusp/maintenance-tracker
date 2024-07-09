package utils

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mailer struct {
	server *mail.SMTPServer
}

func NewMailer() *Mailer {
	server := mail.NewSMTPClient()
	server.Host = "smtp.example.com"
	server.Port = 587
	server.Username = "your-email@example.com"
	server.Password = "your-email-password"
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	return &Mailer{server}
}

func (m *Mailer) SendMail(to, subject, body string) {
	client, err := m.server.Connect()
	if err != nil {
		log.Println("Failed to connect to SMTP server:", err)
		return
	}

	email := mail.NewMSG()
	email.SetFrom("From <your-email@example.com>")
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextPlain, body)

	err = email.Send(client)
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent successfully")
	}
}
