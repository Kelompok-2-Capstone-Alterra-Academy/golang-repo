package service

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmail(to, subject, body string) error {
	from := "refa.developer@gmail.com"
	password := "wmaigneekhhnppec"

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(9999)
	return fmt.Sprintf("%04d", otp)
}
