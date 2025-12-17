package challenge

type Notifier interface {
	Send(message string) error
}
