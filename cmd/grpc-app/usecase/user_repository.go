package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserRepository interface {
	Get(context.Context, string) (domain.User, error)
	UpdateAll(context.Context, domain.User) (domain.User, error)
}
