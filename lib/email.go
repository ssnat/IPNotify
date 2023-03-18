package lib

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

func SendEmail(subject string, content string) error {

	config := GetConfig()

	auth := smtp.PlainAuth(
		"",
		config.Email.SMTPUser,
		config.Email.SMTPPasswd,
		config.Email.SMTPHost,
	)

	to := config.Email.To
	from := mail.Address{Name: "IPNotify", Address: config.Email.From}

	body := fmt.Sprintf(
		"From: " + from.String() + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			content,
	)

	msg := []byte(body)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Email.SMTPHost, config.Email.SMTPPort),
		auth,
		from.Address,
		to,
		msg,
	)

	if err != nil {
		return err
	}

	Logger.Info(fmt.Sprintf("IP information has been sent to email: %s", strings.Join(config.Email.To, ", ")))
	return nil
}

func SendStartupEmail(newIP string) error {
	subject := "Welcome to IPNotify"
	content := fmt.Sprintf(
		"We're happy to inform you that IPNotify is up and running. Whenever your IP address changes, we'll notify you by email with your latest IP address.\n\n"+
			"Current IP: %s\r\n", newIP)
	return SendEmail(subject, content)
}

func SendIPChangeEmail(newIP string, oldIP string) error {
	subject := "Your IP Address Has Changed"
	content := fmt.Sprintf(
		"We are writing to notify you that your IP address has been modified. Below are the details of the change:\n\n"+
			"Previous IP: %s\n"+
			"Current IP: %s\r\n", oldIP, newIP)
	return SendEmail(subject, content)
}
