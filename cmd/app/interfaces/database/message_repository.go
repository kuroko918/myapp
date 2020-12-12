package database

import (
	"github.com/kuroko918/myapp/cmd/app/domain"
)

type MessageRepository struct {
	SqlHandler
}

func (repo *MessageRepository) DeleteById(message domain.Message) (err error) {
	if err = repo.Delete(&message).Error; err != nil {
		return
	}
	return
}

func (repo *MessageRepository) FindAll() (messages domain.Messages, err error) {
	if err = repo.Find(&messages).Error; err != nil {
		return
	}
	return
}

func (repo *MessageRepository) Store(m domain.Message) (message domain.Message, err error) {
	if err = repo.Create(&m).Error; err != nil {
		return
	}
	message = m
	return
}

func (repo *MessageRepository) Update(m domain.Message, attrs ...interface{}) (message domain.Message, err error) {
	if err = repo.Model(&m).Update(attrs).Error; err != nil {
		return
	}
	message = m
	return
}
