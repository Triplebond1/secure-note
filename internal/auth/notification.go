package auth

import "fmt"

type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Sending email:", message)
}

type SMSNotification struct{}

func (s SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}
