package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserRepository interface {
	GetOrStore(context.Context, domain.User) (domain.User, error)
	Get(context.Context, string) (domain.User, error)
}
