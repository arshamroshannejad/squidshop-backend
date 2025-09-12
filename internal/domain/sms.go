package domain

type SmsService interface {
	Send(msg, phone string) error
}
