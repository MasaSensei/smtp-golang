package smtp

import "net/smtp"

const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = 587
	Username   = "username"
	Password   = "password"
)

func SendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", Username, Password, SMTPServer)

	msg := []byte(
		"Subject: " + subject + "\n\n" +
			"From: " + Username + "\n" +
			"To: " + to + "\n\n" +
			"Content-Type: text/plain; charset=utf-8\n\n" +
			body,
	)

	addr := SMTPServer + ":" + SMTPPort

	return smtp.SendMail(addr, auth, Username, []string{to}, msg)
}
