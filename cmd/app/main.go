package main

import (
	"fmt"
	"weekly-newsletter/protocol"
)

// gomail "gopkg.in/gomail.v2"

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		fmt.Println(err)
	}
}

// func testEmail() error {
// 	// https://app.brevo.com/settings/keys/smtp
// 	from := "pramot.nn@gmail.com"
// 	to := "pramot.nn+11@gmail.com"

// 	host := "smtp-relay.sendinblue.com"
// 	port := 587

// 	msg := gomail.NewMessage()
// 	msg.SetHeader("From", from)
// 	msg.SetHeader("To", to)
// 	msg.SetHeader("Subject", "News letter")

// 	// text/html for a html email
// 	msg.SetBody("text/plain", "Welcome to the news letter!")

// 	n := gomail.NewDialer(host, port, from, "Mpy8mxXV1IAs0fJD")

// 	// Send the email
// 	if err := n.DialAndSend(msg); err != nil {
// 		return err
// 	}
// 	return nil
// }
