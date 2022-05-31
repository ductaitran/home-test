package main

type data interface{}

type sender interface {
	send(data)
}

type RESTConfig struct{}

type SMTPConfig struct{}

func (r RESTConfig) send(d data) {
}

func (s SMTPConfig) send(d data) {
}
