package smtp

import (
	"net/smtp"
)

// Konfigurasi SMTP
const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = "587"
	Username   = "your-email@gmail.com"
	Password   = "your-app-password"
)

// SendEmail mengirim email menggunakan SMTP
func SendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", Username, Password, SMTPServer)

	msg := []byte(
		"Subject: " + subject + "\r\n" +
			"From: " + Username + "\r\n" +
			"To: " + to + "\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			body,
	)

	addr := SMTPServer + ":" + SMTPPort
	return smtp.SendMail(addr, auth, Username, []string{to}, msg)
}
