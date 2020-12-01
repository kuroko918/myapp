package usecase

import "github.com/kuroko918/myapp/cmd/app/domain"

type MessageInteractor struct {
	MessageRepository MessageRepository
}

func (interactor *MessageInteractor) Add(u domain.Message) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Store(u)
	return
}

func (interactor *MessageInteractor) DeleteById(u domain.Message) (err error) {
	err = interactor.MessageRepository.DeleteById(u)
	return
}

func (interactor *MessageInteractor) Messages() (messages domain.Messages, err error) {
	messages, err = interactor.MessageRepository.FindAll()
	return
}

func (interactor *MessageInteractor) Update(u domain.Message) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Update(u)
	return
}
