package controllers

import (
	"github.com/kuroko918/myapp/cmd/app/domain"
	"github.com/kuroko918/myapp/cmd/app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/app/usecase"
)

type MessagesController struct {
	Interactor usecase.MessageInteractor
}

func NewMessagesController(sqlHandler database.SqlHandler) *MessagesController {
	return &MessagesController{
		Interactor: usecase.MessageInteractor{
			MessageRepository: &database.MessageRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *MessagesController) Create(c Context) {
	u := domain.Message{}
	c.Bind(&u)
	message, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, message)
}

func (controller *MessagesController) Delete(c Context) {
	id := c.Param("id")
	message := domain.Message{
		ID: id,
	}
	err := controller.Interactor.DeleteById(message)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, message)
}

func (controller *MessagesController) Index(c Context) {
	messages, err := controller.Interactor.Messages()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, messages)
}

func (controller *MessagesController) Update(c Context) {
	u := domain.Message{}
	c.Bind(&u)
	message, err := controller.Interactor.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, message)
}
