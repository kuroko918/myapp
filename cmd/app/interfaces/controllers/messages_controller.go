package controllers

import (
	"strconv"

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
	userId, _ := c.Get("AuthenticatedUserId")
	message := domain.Message{
		UserId: userId.(string),
	}
	c.Bind(&message)
	message, err := controller.Interactor.Add(message)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, message)
}

func (controller *MessagesController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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
	id, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")
	message := domain.Message{
		ID: id,
	}
	message, err := controller.Interactor.Update(message, "Content", content)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, message)
}
