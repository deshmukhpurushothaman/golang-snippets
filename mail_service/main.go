package main

import (
	"fmt"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func sendSimpleMail(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"deshmukhpurushothaman@gmail.com",
		"oyiuuiykoiemcopo",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"deshmukhpurushothaman@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendSimpleHTMLMail(subject string, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"deshmukhpurushothaman@gmail.com",
		"oyiuuiykoiemcopo",
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"deshmukhpurushothaman@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendGoMail(subject string, html string, to []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "deshmukhpurushothaman@gmail.com")
	m.SetHeader("To", "deshmukhpurushothaman@gmail.com", "deshmukhpurushothaman@gmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "deshmukhpurushothaman@gmail.com", "oyiuuiykoiemcopo")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	// sendSimpleMail("From argument", "Changed to arguments", []string{"deshmukhpurushothaman@gmail.com"})
	// sendSimpleHTMLMail(
	// 	"From argument",
	// 	"<h1>HTML header tag</h1><p>Email as html template</p>",
	// 	[]string{"deshmukhpurushothaman@gmail.com"},
	// )
	sendGoMail(
		"From argument",
		"<h1>HTML header tag</h1><p>Email as html template</p>",
		[]string{"deshmukhpurushothaman@gmail.com"},
	)
}
