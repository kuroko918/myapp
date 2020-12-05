package domain

import "time"

type Message struct {
	Id        int32
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Messages []Message
