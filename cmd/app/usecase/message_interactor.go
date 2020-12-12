package usecase

import "github.com/kuroko918/myapp/cmd/app/domain"

type MessageInteractor struct {
	MessageRepository MessageRepository
}

func (interactor *MessageInteractor) Add(m domain.Message) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Store(m)
	return
}

func (interactor *MessageInteractor) DeleteById(m domain.Message) (err error) {
	err = interactor.MessageRepository.DeleteById(m)
	return
}

func (interactor *MessageInteractor) Messages() (messages domain.Messages, err error) {
	messages, err = interactor.MessageRepository.FindAll()
	return
}

func (interactor *MessageInteractor) Update(m domain.Message, attrs ...interface{}) (message domain.Message, err error) {
	message, err = interactor.MessageRepository.Update(m)
	return
}
