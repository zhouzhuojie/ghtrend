package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var (
	subject  = "Github Trending Daily"
	smtpHost = "smtp.gmail.com"
	smtpPort = 587

	// auth should be loaded from env or config
	auth = smtp.PlainAuth("", userEmail, userPassword, smtpHost)
)

// SendMail sends the email
func SendMail(e *email.Email) error {
	return e.Send(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth)
}

// SendGithubTrendMail sends the github trending email
func SendGithubTrendMail(html []byte) error {
	e := &email.Email{
		From:    userEmail,
		To:      []string{userEmail},
		Subject: subject,
		HTML:    html,
	}
	return SendMail(e)
}
