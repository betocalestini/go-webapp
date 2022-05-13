package utils

type Alert struct {
	Message string
	Type    string
	Active  string
}

func NewAlert(message, alert, active string) Alert {
	return Alert{message, alert, active}
}
