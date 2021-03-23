package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type MessageRepository interface {
	DeleteById(context.Context, domain.Message) error
	Get(context.Context, string) (domain.Message, error)
	FindAll(context.Context) ([]domain.Message, error)
	Store(context.Context, domain.Message) (domain.Message, error)
	Update(context.Context, map[string]interface{}) (domain.Message, error)
}
