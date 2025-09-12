package domain

type Repository interface {
	User() UserRepository
}
