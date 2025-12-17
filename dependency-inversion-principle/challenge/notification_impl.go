package challenge

import "fmt"

type NotificationManager struct {
	Service Notifier
}

func (n *NotificationManager) BlastNotification(msg string) {
	err := n.Service.Send(msg)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
