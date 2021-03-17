package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UpdateAll(ctx context.Context, u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.UpdateAll(ctx, u)
	return
}

func (interactor *UserInteractor) User(ctx context.Context, userId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.Get(ctx, userId)
	return
}
