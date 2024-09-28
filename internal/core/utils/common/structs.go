package common

type Module struct {
	Patterns []string
	Handler  MessageHandler
}

type ActionLog struct {
	Request  any
	Response any
}
