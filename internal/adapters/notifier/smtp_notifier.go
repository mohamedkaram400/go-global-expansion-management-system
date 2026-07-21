package notifier

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
)

type SMTPNotifier struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewSMTPNotifier(host string, port int, user, pass, from string) *SMTPNotifier {
	return &SMTPNotifier{
		Host:     host,
		Port:     port,
		Username: user,
		Password: pass,
		From:     from,
	}
}

func (s *SMTPNotifier) SendMatchNotification(ctx context.Context, payload ports.MatchNotification) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	msg := []byte("From: " + s.From + "\r\n" +
		"To: " + strings.Join(payload.To, ",") + "\r\n" +
		"Subject: " + payload.Subject + "\r\n\r\n" +
		payload.Body + "\r\n")

	// fmt.Println("msg: ", msg, "addr: ", addr, "auth: ", auth)

	err := smtp.SendMail(addr, auth, s.From, payload.To, msg)
	if err != nil {
		log.Printf("SMTP send error: %v", err)
		return err
	}
	log.Println(">>> email sent successfully to", payload.To)
	return nil
}
