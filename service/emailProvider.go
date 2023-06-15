package service

type emailProvider interface {
	sendEmail() error
}
