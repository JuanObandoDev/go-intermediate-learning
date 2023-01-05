package main

import "fmt"

// creating interfaces
type INorificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// creating a factory
type SMSNotification struct{}

func (smsn SMSNotification) SendNotification() {
	fmt.Println("Sending SMS Notification")
}

func (smsn SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// creating a sender
type SMSNotificationSender struct{}

func (smsns SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (smsns SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// creating a factory
type EmailNotification struct{}

func (en EmailNotification) SendNotification() {
	fmt.Println("Sending Email Notification")
}

func (en EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// creating a sender
type EmailNotificationSender struct{}

func (ens EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (ens EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}
