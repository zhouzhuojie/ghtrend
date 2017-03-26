package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

var (
	userEmail    = os.Getenv("userEmail")
	userPassword = os.Getenv("userPassword")
	smtpHost     = "smtp.gmail.com"
	smtpPort     = 587
	subject      = "Github Trending Daily"

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
