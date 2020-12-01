package domain

import "time"

type Message struct {
	ID        string
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Messages []Message
