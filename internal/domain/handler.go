package domain

type Handler interface {
	User() UserHandler
}
