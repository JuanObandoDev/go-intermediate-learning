package main

import (
	"fmt"
	"strings"
)

// creating interfaces
type INotificationFactory interface {
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

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	notificationType = strings.ToLower(notificationType)
	if notificationType == "sms" {
		return &SMSNotification{}, nil
	} else if notificationType == "email" {
		return &EmailNotification{}, nil
	} else {
		return nil, fmt.Errorf("No notification type found")
	}
}

func sendNotification(inf INotificationFactory) {
	inf.SendNotification()
}

func getMethod(inf INotificationFactory) {
	fmt.Println(inf.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emaiFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emaiFactory)

	getMethod(smsFactory)
	getMethod(emaiFactory)
}
