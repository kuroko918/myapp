package usecase

import "github.com/kuroko918/myapp/cmd/grpc-app/domain"

type UserRepository interface {
	GetOrStore(domain.User) (domain.User, error)
}
