package smtp

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func InitSMTP() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Failed to load .env file")
	} else {
		log.Println("Successfully loaded .env file")
	}
}

func SendEmail(to string, subject string, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Failed to load .env file")
		return err
	} else {
		log.Println("Successfully loaded .env file")
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", username, password, smtpServer)

	msg := []byte(
		"Subject: " + subject + "\r\n" +
			"From: " + username + "\r\n" +
			"To: " + to + "\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			body,
	)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpServer, smtpPort), tlsconfig)
	if err != nil {
		log.Println("Error connecting to SMTP server:", err)
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		log.Println("Error creating SMTP client:", err)
		return err
	}

	if err = c.Auth(auth); err != nil {
		log.Println("Error authenticating:", err)
		return err
	}

	if err = c.Mail(username); err != nil {
		log.Println("Error sending MAIL command:", err)
		return err
	}

	if err = c.Rcpt(to); err != nil {
		log.Println("Error sending RCPT command:", err)
		return err
	}

	w, err := c.Data()
	if err != nil {
		log.Println("Error starting data transfer:", err)
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		log.Println("Error writing data:", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Println("Error closing data:", err)
		return err
	}

	err = c.Quit()
	if err != nil {
		log.Println("Error quitting SMTP session:", err)
		return err
	}

	return nil
}
