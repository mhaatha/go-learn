package main

import (
	"fmt"

	"github.com/mhaatha/go-learn/dependency-inversion-principle/challenge"
)

func main() {
	emailManager := &challenge.NotificationManager{
		Service: &challenge.EmailService{},
	}
	emailManager.BlastNotification("Halo via Email!")

	smsManager := &challenge.NotificationManager{
		Service: &challenge.SMSService{},
	}
	smsManager.BlastNotification("Halo via SMS!")

	fmt.Println("\n--- Switching Runtime ---")
	smsManager.Service = &challenge.EmailService{}
	smsManager.BlastNotification("Sekarang jadi Email lagi!")
}
