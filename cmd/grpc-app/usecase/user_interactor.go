package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) User(ctx context.Context, userId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.Get(ctx, userId)
	return
}

func (interactor *UserInteractor) GetOrAdd(ctx context.Context, u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.GetOrStore(ctx, u)
	return
}
