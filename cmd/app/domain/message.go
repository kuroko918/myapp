package domain

import "time"

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Messages []Message
