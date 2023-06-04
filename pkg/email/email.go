package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

type service struct {
	emailPort    int
	emailHost    string
	emailAddress string
	emailAppPass string
}

type EmailIntf interface {
	SendEmail(receiverEmail string, subject string, messageBody string) (bool, error)
}

// NewEmailService() - initializes email service
func NewEmailService(emailPort int, emailHost string, emailAddress string, emailAppPass string) EmailIntf {
	svc := service{}
	svc.emailPort = emailPort
	svc.emailHost = emailHost
	svc.emailAddress = emailAddress
	svc.emailAppPass = emailAppPass
	return &svc
}

// SendEmail() - sends email by passing receiver email, subject and the message body
func (s *service) SendEmail(receiverEmail string, subject string, messageBody string) (bool, error) {
	isSend := false
	emailMessage := s.buildEmailMessage(receiverEmail, subject, messageBody)
	if emailMessage == "" {
		return isSend, errors.New("empty")
	}
	// Create authentication for the sender's email
	auth := smtp.PlainAuth("", s.emailAddress, s.emailAppPass, s.emailHost)

	// Connect to the SMTP server
	smtpServer := fmt.Sprintf("%s:%d", s.emailHost, s.emailPort)
	conn, err := smtp.Dial(smtpServer)
	if err != nil {
		return isSend, err
	}
	defer conn.Close()

	// Start TLS encryption
	tlsConfig := &tls.Config{
		ServerName: s.emailHost,
	}
	if err := conn.StartTLS(tlsConfig); err != nil {
		log.Fatal(err)
	}

	// Authenticate with the SMTP server
	if err := conn.Auth(auth); err != nil {
		return isSend, err
	}

	// Set the sender and recipient
	if err := conn.Mail(s.emailAddress); err != nil {
		return isSend, err
	}
	if err := conn.Rcpt(receiverEmail); err != nil {
		return isSend, err
	}

	// Send the email message
	writer, err := conn.Data()
	if err != nil {
		return isSend, err
	}
	_, err = writer.Write([]byte(emailMessage))
	if err != nil {
		return isSend, err
	}
	err = writer.Close()
	if err != nil {
		return isSend, err
	}

	isSend = true

	return isSend, nil
}

func (s *service) buildEmailMessage(receiverEmail string, subject string, messageBody string) string {
	if receiverEmail == "" || subject == "" || messageBody == "" || s.emailAddress == "" {
		return ""
	}
	emailMessage := "From: " + s.emailAddress + "\r\n" +
		"To: " + receiverEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		messageBody

	return emailMessage
}
