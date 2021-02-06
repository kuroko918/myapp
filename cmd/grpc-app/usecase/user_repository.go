package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserRepository interface {
	Get(context.Context, string) (domain.User, error)
}
