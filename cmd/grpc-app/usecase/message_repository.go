package usecase

import "github.com/kuroko918/myapp/cmd/grpc-app/domain"

type MessageRepository interface {
	DeleteById(domain.Message) error
	FindAll() ([]domain.Message, error)
	Store(domain.Message) (domain.Message, error)
	Update(domain.Message, ...interface{}) (domain.Message, error)
}
