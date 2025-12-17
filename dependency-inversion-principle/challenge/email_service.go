package challenge

import "fmt"

type EmailService struct{}

func (e *EmailService) Send(message string) error {
	fmt.Println("Sending Email:", message)
	return nil
}
