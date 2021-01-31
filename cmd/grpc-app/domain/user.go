package domain

import "time"

type User struct {
	ID        string `json:"id"`
  Name      string `json:"name"`
  Email     string `json:"email"`
  Avatar    string `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
