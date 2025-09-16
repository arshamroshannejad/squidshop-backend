package domain

type Handler interface {
	Auth() AuthHandler
	User() UserHandler
	Category() CategoryHandler
}
