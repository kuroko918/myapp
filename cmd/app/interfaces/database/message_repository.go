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

func (repo *MessageRepository) FindAll() (messages []domain.Message, err error) {
	if err = repo.Preload("User").Find(&messages).Error; err != nil {
		return
	}
	return
}

func (repo *MessageRepository) Store(m domain.Message) (message domain.Message, err error) {
	if err = repo.Create(&m).Preload("User").Find(&m).Error; err != nil {
		return
	}
	message = m
	return
}

func (repo *MessageRepository) Update(m domain.Message, attrs ...interface{}) (message domain.Message, err error) {
	if err = repo.Preload("User").First(&m).Update(attrs).Error; err != nil {
		return
	}
	message = m
	return
}
