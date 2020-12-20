package usecase

import "github.com/kuroko918/myapp/cmd/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetOrAdd(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.GetOrStore(u)
	return
}
