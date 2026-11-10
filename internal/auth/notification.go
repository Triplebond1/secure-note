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

type PushNotification struct{}

type Notifier interface {
	Send(message string)
	Close()
}

func (e EmailNotification) Close() {
	fmt.Println("Email notification closed")
}

func (s SMSNotification) Close() {
	fmt.Println("SMS notification closed")
}

type MessageSender interface {
	Send(message string)
}

type MessageCloser interface {
	MessageSender
	Close()
}

type Anything interface{}
