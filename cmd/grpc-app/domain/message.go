package domain

import "time"

type Message struct {
	ID        uint64    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	User      User `json:"user"`
}
