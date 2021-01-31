package usecase

import (
	"context"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type MessageInteractor struct {
	MessageRepository MessageRepository
}

func (interactor *MessageInteractor) Add(ctx context.Context, m domain.Message) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Store(ctx, m)
	return
}

func (interactor *MessageInteractor) DeleteById(ctx context.Context, m domain.Message) (err error) {
	err = interactor.MessageRepository.DeleteById(ctx, m)
	return
}

func (interactor *MessageInteractor) Messages(ctx context.Context) (messages []domain.Message, err error) {
	messages, err = interactor.MessageRepository.FindAll(ctx)
	return
}

func (interactor *MessageInteractor) Update(ctx context.Context, m map[string]interface{}) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Update(ctx, m)
	return
}
