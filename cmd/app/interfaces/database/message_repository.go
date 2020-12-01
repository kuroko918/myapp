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

func (repo *MessageRepository) Store(u domain.Message) (message domain.Message, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	message = u
	return
}

func (repo *MessageRepository) Update(u domain.Message) (message domain.Message, err error) {
	if err = repo.Save(&u).Error; err != nil {
		return
	}
	message = u
	return
}
