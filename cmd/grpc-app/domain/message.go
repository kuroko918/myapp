package domain

import "time"

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User `json:"user"`
}
