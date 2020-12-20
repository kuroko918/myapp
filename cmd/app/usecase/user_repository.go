package usecase

import "github.com/kuroko918/myapp/cmd/app/domain"

type UserRepository interface {
	GetOrStore(domain.User) (domain.User, error)
}
