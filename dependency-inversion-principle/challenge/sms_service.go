package challenge

import "fmt"

type SMSService struct{}

func (s *SMSService) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}
