package lib

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

func SendEmail(newIp string, oldIp string) error {

	config := GetConfig()

	auth := smtp.PlainAuth(
		"",
		config.Email.SMTPUser,
		config.Email.SMTPPasswd,
		config.Email.SMTPHost,
	)

	to := config.Email.To
	from := mail.Address{Name: "IPNotify Service", Address: config.Email.From}

	body := fmt.Sprintf(
		"From: "+from.String()+"\r\n"+
			"Subject: [Important] Your IP Address Has Changed\r\n"+
			"\r\n"+
			"We are writing to notify you that your IP address has been modified. Below are the details of the change:\n\n"+
			"Previous IP: [%s]\n"+
			"New IP: [%s]\r\n", oldIp, newIp)

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
